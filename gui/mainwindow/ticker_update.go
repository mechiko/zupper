package mainwindow

import (
	"time"
	"zupper/domain"
)

// каждый тик проверяем канал на входящие сообщения
func (w *MainWindow) StartTicker(period time.Duration) {
	w.ticker = time.NewTicker(period)
	go func() {
		for range w.ticker.C {
			w.tick()
		}
	}()
}

func (w *MainWindow) StopTicker() {
	if w.ticker != nil {
		w.ticker.Stop()
	}
}

// эту функцию передаем в качесте callback для работы с каналом
func (w *MainWindow) SendChanel(m domain.Model) {
	select {
	case w.InChangeModel <- m:
		// Message sent successfully
	default:
		w.Logger().Warnf("Channel buffer full, dropping model update: %s", m)
	}
}

// каждые N миллисекунд проверяем канал уведомлений смены модели
func (w *MainWindow) tick() {
	select {
	case m := <-w.InChangeModel:
		w.Logger().Debugf("chanel receive model chang state %s", m)
		switch m {
		case domain.StatusBar:
			w.Synchronize(func() {
				w.UpdateStatusBar()
			})
		default:
			w.Synchronize(func() {
				page := w.Tvm.CurrentPage()
				if page == nil {
					w.Logger().Warn("tick: no current page to update")
					return
				}
				page.Update()
			})
		}
	default:
	}
}
