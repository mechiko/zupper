package mainwindow

import (
	"fmt"

	"zupper/gui/types"

	"github.com/mechiko/walk"
)

// вызывается при смене страницы
func (w *MainWindow) changePage() error {
	currentItem := w.tv.CurrentItem()
	menu, ok := currentItem.(*types.AppMenu)
	if !ok {
		return fmt.Errorf("mainwindows changePage: unexpected menu type %T", currentItem)
	}
	// если action для пункта меню не установлено, то не меняем страницу и пункт меню
	if menu.Action() == nil {
		return nil
	}
	if err := w.SetCurrentMenu(menu); err != nil {
		return fmt.Errorf("mainwindows SetCurrentMenu %w", err)
	}
	return nil
}

func (w *MainWindow) CurrentPageTitle() string {
	if w.Tvm.CurrentMenu == nil {
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
	}

	if prevPage != nil {
		w.pageCom.SaveState()
		prevPage.SetVisible(false)
		if widget, ok := prevPage.(walk.Widget); ok {
			widget.SetParent(nil)
		}
		prevPage.Disposing().Attach(func() {
		})
		prevPage.Dispose()
		prevPage.Clear()
	}

	newPage, exists := w.Tvm.Menu2NewPage[pageMenu]
	if !exists {
		return fmt.Errorf("no page factory found for menu: %s", pageMenu.Name())
	}

	if w.pageCom.Children().Len() > 0 {
		w.DisposeChildren(w.pageCom)
	}

	page, err := newPage(w.pageCom, w.Apper, w.repo)
	if err != nil {
		return fmt.Errorf("newPage(w.pageCom) %w", err)
	}
	// через эту функцию прописываем метод отправки в канал смены состояния
	page.SetSendFunc(w.SendChanel)

	w.Tvm.SetCurrentPage(page)
	w.Tvm.CurrentMenu = pageMenu
	// создаем событие смены страницы
	w.currentPageChangedPublisher.Publish()
	// w.SetFocus()
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
