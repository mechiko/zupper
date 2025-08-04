package gui

import (
	"context"
	"fmt"
	"time"
	"zupper/reductor"
)

var ticketStateChange = time.Millisecond * 20

func (s *guiService) Run(ctx context.Context) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("gui Run panic %v", r)
		}
	}()

	go func() {
		// если получен по контексту сигнал завершения то вызываем завершение приложения
		<-ctx.Done()
		s.Logger().Info("gui:run receive <-ctx.Done() and attempt launch windows.close()")
		// завершение идет глубже, есть на событие закрытие окна метод в службе там завершаем приложение walk
		// в s.Shutdown() только закрытие окна майн
		// вызов завершения приложения здесь не дает результата из Run не выходи, а если в событии вызывать выходит
		if err := s.Shutdown(); err != nil {
			s.Logger().Errorf("gui:run shutdown %v", err)
		}
		s.Logger().Info("gui:run after shutdown")
	}()
	reductor.Instance().SetOutChanState(s.MainWindow.InChangeModel)
	s.MainWindow.Starting().Attach(func() {
		// здесь можно что то сделать при запуске главного окна
		go s.MainWindow.StartTicker(ticketStateChange)
	})

	if codeExit := s.MainWindow.Run(); codeExit != 0 {
		s.Logger().Infof("gui:walk exit from MainWindow.Run() CODE=%v", codeExit)
		return fmt.Errorf("GUI error")
	} else {
		s.Logger().Infof("gui:walk exit from MainWindow.Run() CODE=%v", codeExit)
		return nil
	}
}
