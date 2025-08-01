package mainwindow

import (
	"fmt"

	"zupper/entity"
	"zupper/gui/types"

	"github.com/mechiko/walk"
)

// вызывается при смене страницы
func (w *MainWindow) сhangePage() error {
	menu := w.tv.CurrentItem().(*types.AppMenu)
	if menu.Action() == nil {
		return nil
	}
	return w.SetCurrentMenu(menu)
}

func (w *MainWindow) CurrentPageTitle() string {
	if w.Tvm.CurrentPage() == nil {
		return ""
	}
	return w.Tvm.CurrentMenu.Name()
}

func (w *MainWindow) CurrentPageChanged() *walk.Event {
	return w.currentPageChangedPublisher.Event()
}

func (w *MainWindow) SetCurrentMenu(pageMenu *types.AppMenu) error {
	defer func() {
		// if !w.pageCom.IsDisposed() {
		// 	w.pageCom.RestoreState()
		// }
		if r := recover(); r != nil {
			panic(fmt.Errorf("%s SetCurrentMenu panic %v", modError, r))
		}
	}()

	prevPage := w.Tvm.CurrentPage()

	if prevPage == nil {
		// еще меню не выставилось значит первый запуск и активация первого меню
		msg := entity.Message{
			Sender: "mainwindow.SetCurrentMenu",
			Cmd:    "first",
			Model:  nil,
		}
		w.IApp.Effects().ChanIn() <- msg
	}

	if prevPage != nil {
		w.pageCom.SaveState()
		prevPage.SetVisible(false)
		prevPage.(walk.Widget).SetParent(nil)
		prevPage.Disposing().Attach(func() {
		})
		prevPage.Dispose()
		prevPage.Clear()
	}

	newPage := w.Tvm.Menu2NewPage[pageMenu]

	if w.pageCom.Children().Len() > 0 {
		w.DisposeChildren(w.pageCom)
	}

	page, err := newPage(w.pageCom, w.IApp)
	if err != nil {
		return fmt.Errorf("newPage(w.pageCom) %w", err)
	}

	w.Tvm.SetCurrentPage(page)
	w.Tvm.CurrentMenu = pageMenu
	// через редуктор вызываем обновление состояния GUI
	// запрашиваем текущее состояние в редукторе которое прилетит в обновление
	// каждый вид пересоздается при активации поэтому надо состояние присылать
	msg := entity.Message{
		Sender: "mainwindow.SetCurrentMenu",
		Cmd:    pageMenu.Class(),
		Model:  nil,
	}
	w.IApp.Effects().ChanIn() <- msg
	// w.Logger().Debugf("%s reductor chanin msg %s", modError, msg.Cmd)
	// msg = entity.Message{
	// 	Sender: "mainwindow.SetCurrentMenu",
	// 	Cmd:    "license",
	// 	Model:  nil,
	// }
	// w.IApp.Effects().ChanIn() <- msg
	// w.Logger().Debugf("%s effects chanin msg %s", modError, msg.Cmd)

	// создаем событие смены страницы
	w.currentPageChangedPublisher.Publish()

	w.SetFocus()

	return nil
}

// func (w *MainWindow) CheckChildren(wtest walk.Container) error {
// 	return nil
// }

func (w *MainWindow) DisposeChildren(wtest walk.Container) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s DisposeChildren panic %v", modError, r)
		}
	}()

	ws := wtest.Children().Len()
	for i := 0; i < ws; i++ {
		w := wtest.Children().At(i)
		w.Dispose()
	}
	return nil
}
