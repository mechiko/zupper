package mainwindow

import (
	"zupper/domain"
	"zupper/domain/models/statusbar"
	"zupper/reductor"
)

func (w *MainWindow) UpdateStatusBar() {
	modelReductor, err := reductor.Instance().Model(domain.StatusBar)
	if err != nil {
		w.Logger().Errorf("gui:statusbar update %v", err)
		return
	}
	model, ok := modelReductor.(*statusbar.StatusBar)
	if !ok {
		w.Logger().Errorf("gui:statusbar wrong model reductor %T", modelReductor)
		return
	}
	if model.Utm {
		w.SbiUtmState.SetText("УТМ")
		w.SbiUtmState.SetIcon(w.IconGreen)
	} else {
		w.SbiUtmState.SetText("УТМ")
		w.SbiUtmState.SetIcon(w.IconRed)
	}
	if model.Scan {
		w.SbiScan.SetText("Прием")
		w.SbiScan.SetIcon(w.IconGreen)
	} else {
		w.SbiScan.SetText("Прием")
		w.SbiScan.SetIcon(w.IconRed)
	}
	if model.License {
		w.SbiLicense.SetText("Лицензия")
		w.SbiLicense.SetIcon(w.IconGreen)
	} else {
		w.SbiLicense.SetText("Лицензия")
		w.SbiLicense.SetIcon(w.IconRed)
	}
	w.SbiFsrarId.SetText(model.FsrarID)
}
