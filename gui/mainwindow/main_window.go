package mainwindow

import (
	"fmt"
	"image/color"

	"zupper/gui/resource"
)

func (w *MainWindow) Create(cfg *MainWindowConfig) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s Create %v", modError, r)
		}
	}()

	if svgIcon, err := resource.New(w).Svg(resource.SvgCircle, color.RGBA{R: 255, A: 255}, 18, 18); err != nil {
		w.Logger().Errorf("w:resource error %s", err)
	} else {
		w.IconRed = svgIcon
	}
	if svgIcon, err := resource.New(w).Svg(resource.SvgCircle, color.RGBA{G: 255, A: 255}, 18, 18); err != nil {
		w.Logger().Errorf("w:resource error %s", err)
	} else {
		w.IconGreen = svgIcon
	}

	if err := w.dclCreate(cfg); err != nil {
		return err
	}
	succeeded := false
	defer func() {
		if !succeeded {
			w.Dispose()
		}
	}()
	fontMain := w.Font()
	if fontMain != nil {
		w.Logger().Errorf("gui:mainwindow font %v", err)
	} else {
		w.StatusBar().SetFont(fontMain)
	}

	if w.Tvm != nil {
		w.tv.SetCurrentItem(w.Tvm.DefaultMenu())
	} else {
		return fmt.Errorf("mainwindow AppmenuTreeModel is nil")
	}
	w.CurrentPageChanged().Attach(cfg.OnCurrentPageChanged)
	succeeded = true
	w.changePage()
	// if recover block present then return err
	return err
}
