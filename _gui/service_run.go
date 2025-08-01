package gui

import (
	"context"
	"fmt"

	"zupper/entity"
)

func (s *guiService) Run(ctx context.Context) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s Run panic %v", modError, r)
		}
	}()

	go func() {
		// если получен по контексту сигнал завершения то вызываем завершение приложения
		<-ctx.Done()
		s.Logger().Info("gui:run receive <-ctx.Done() launch windows.close()")
		// завершение идет глубже, есть на событие закрытие окна метод в службе там завершаем приложение walk
		// в s.Shutdown() только закрытие окна майн
		// вызов завершения приложения здесь не дает результата из Run не выходи, а если в событии вызывать выходит
		s.Shutdown()
		s.Logger().Info("gui:run receive <-ctx.Done()")
	}()
	s.MainWindow.Starting().Attach(func() {
		// здесь можно что то сделать при запуске главного окна
	})

	if ii := s.MainWindow.Run(); ii != 0 {
		s.Logger().Infof("gui:walk exit from MainWindow.Run() CODE=%v", ii)
		return fmt.Errorf("GUI error")
	} else {
		s.Logger().Infof("gui:walk exit from MainWindow.Run() CODE=%v", ii)
	}
	s.Logger().Infof("gui:walk exit entity.ErrAppShutdown")
	return entity.ErrAppShutdown
}
