package mainwindow

import (
	"fmt"
	"time"

	"zupper/domain"
	"zupper/gui/types"

	"github.com/mechiko/walk"
)

// вызывается при смене страницы
func (w *MainWindow) changePage() error {
	menu := w.tv.CurrentItem().(*types.AppMenu)
	if menu.Action() == nil {
		return nil
	}
	if err := w.SetCurrentMenu(menu); err != nil {
		return fmt.Errorf("mainwindows SetCurrentMenu %w", err)
	}
	return nil
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

	page, err := newPage(w.pageCom, w.Apper, w.repo)
	if err != nil {
		return fmt.Errorf("newPage(w.pageCom) %w", err)
	}
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

// каждый тик проверяем канал на входящие сообщения
func (w *MainWindow) StartTicker(period time.Duration) {
	w.ticker = time.NewTicker(period)
	for range w.ticker.C {
		w.tick()
	}
}

func (w *MainWindow) StopTicker() {
	if w.ticker != nil {
		w.ticker.Stop()
	}
}

func (w *MainWindow) SendChanel(m domain.Model) {
	w.InChangeModel <- m
}

func (w *MainWindow) tick() {
	select {
	case m := <-w.InChangeModel:
		w.Logger().Debugf("chanel receive model chang state %s", m)
		switch m {
		case domain.StatusBar:
		default:

		}
		w.Synchronize(func() {
			page := w.Tvm.CurrentPage()
			page.Update()
		})
	default:
	}
}
