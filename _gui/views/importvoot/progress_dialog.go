package importvoot

import (
	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"
)

func (p *ImportTTNPage) progressDialog(f func(*walk.Dialog, *walk.ProgressBar)) {
	defer func() {
		if r := recover(); r != nil {
			p.app.Logger().Errorf("gui:dlg:progress panic %v", r)
		}
	}()
	var dlg *walk.Dialog
	var pb *walk.ProgressBar

	if err := (dcl.Dialog{
		AssignTo:   &dlg,
		Title:      "Обработка",
		Borderless: true,
		Layout: dcl.VBox{
			Margins:     dcl.Margins{Left: 5, Right: 5, Top: 0, Bottom: 5},
			SpacingZero: true,
			Alignment:   dcl.Alignment2D(dcl.AlignHNearVNear),
		},
		Children: []dcl.Widget{
			dcl.ProgressBar{
				Name:       "pb",
				AssignTo:   &pb,
				MinSize:    dcl.Size{Width: 200, Height: 30},
				Background: dcl.SolidColorBrush{Color: walk.RGB(0, 155, 155)},
				MaxValue:   100,
				MinValue:   1,
				Value:      1,
			},
		},
	}).Create(p.form); err != nil {
		p.app.Logger().Errorf("gui:dlg:progress create slg %s", err.Error())
	}
	dlg.Starting().Attach(func() {
		go f(dlg, pb)
	})
	dlg.Run()
}
