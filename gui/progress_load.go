package gui

import (
	"fmt"
	"time"

	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"
)

// func (f *ChronicalFilter) Submit() {
// 	// если использую DataBinder и на структуру по ссылке то он сам обновляет
// 	core.App().Logger().Debug("Submit ChronicalFilter")
// }

func (s *guiService) ProgressDialog(pw walk.Form, cb func() int) (ii int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s panic %v", modError, r)
		}
	}()
	var dlg *walk.Dialog
	var pb *walk.ProgressBar

	if err := (dcl.Dialog{
		AssignTo: &dlg,
		Visible:  false,
		Title:    "Обработка",
		// MinSize:    dcl.Size{Width: 200, Height: 70},
		Borderless: true,
		// Layout:  VBox{MarginsZero: true, SpacingZero: true, Alignment: Alignment2D(AlignHNearVNear)},
		Layout: dcl.VBox{
			Margins:     dcl.Margins{Left: 5, Right: 5, Top: 0, Bottom: 5},
			SpacingZero: true,
			Alignment:   dcl.Alignment2D(dcl.AlignHNearVNear),
		},
		Children: []dcl.Widget{
			// dcl.Label{
			// 	Text:      "Экспорт в файл",
			// 	Alignment: dcl.Alignment2D(dcl.AlignCenter),
			// },
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
	}).Create(pw); err != nil {
		return 0, err
	}
	// x := int((pw.Width() - dlg.Size().Width) / 2)
	// y := int((pw.Height() - dlg.Size().Height) / 2)
	// dlg.SetVisible(true)
	// fmt.Printf("mainwindow bounds %v\n", pw.BoundsPixels())
	// fmt.Printf("dlg before set bounds %v\n", dlg.BoundsPixels())
	// dlg.SetBoundsPixels(walk.Rectangle{
	// 	X:      x - 500,
	// 	Y:      y - 500,
	// 	Width:  dlg.Size().Width,
	// 	Height: dlg.Size().Height,
	// })
	// fmt.Printf("dlg after set bounds %v\n", dlg.BoundsPixels())
	// каждые 100 мс вызываем функцию переданную в диалог и по ней выходим когда вернет 100
	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		for range ticker.C {
			step := cb()
			pb.SetValue(step)
			if step > 100 {
				ticker.Stop()
				dlg.Cancel()
			}
		}
	}()

	// второй вариант
	// когда вызываем внешнюю обработку и передаем туда функцию для счетчика прогресса
	// err := p.ExportExcelOoxmlCb(func(v int) {
	// 	if v > 100 {
	// 		return
	// 	}
	// 	pb.SetValue(v)
	// })

	ret := dlg.Run()
	dlg.Dispose()
	return ret, nil
}
