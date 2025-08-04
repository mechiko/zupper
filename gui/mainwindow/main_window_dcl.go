package mainwindow

import (
	"fmt"

	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"
)

func (w *MainWindow) dclCreate(cfg *MainWindowConfig) error {
	if err := (dcl.MainWindow{
		AssignTo: &w.MainWindow,
		Name:     cfg.Name,
		Title:    cfg.Title,
		Enabled:  cfg.Enabled,
		// Visible:  false,
		Font:    cfg.Font,
		MinSize: cfg.MinSize,
		// MaxSize:          cfg.MaxSize,
		// Size:             dcl.Size{Width: 1000, Height: 700},
		Size:             cfg.MinSize,
		MenuItems:        cfg.MenuItems,
		ToolBar:          cfg.ToolBar,
		ContextMenuItems: cfg.ContextMenuItems,
		OnKeyDown:        cfg.OnKeyDown,
		OnKeyPress:       cfg.OnKeyPress,
		OnKeyUp:          cfg.OnKeyUp,
		OnMouseDown:      cfg.OnMouseDown,
		OnMouseMove:      cfg.OnMouseMove,
		OnMouseUp:        cfg.OnMouseUp,
		Layout:           dcl.HBox{MarginsZero: true, SpacingZero: true},
		Children: []dcl.Widget{
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: true},
				Children: []dcl.Widget{
					dcl.HSplitter{
						Children: []dcl.Widget{
							dcl.Composite{
								Layout: dcl.HBox{MarginsZero: true, SpacingZero: true,
									Margins: dcl.Margins{Left: 5, Top: 0, Right: 0, Bottom: 0}},
								Children: []dcl.Widget{
									dcl.TreeView{
										AssignTo:      &w.tv,
										StretchFactor: 1,
										Model:         w.Tvm,
										OnCurrentItemChanged: func() {
											defer func() {
												if recoverVar := recover(); recoverVar != nil {
													w.Logger().Errorf("mainwindow tv change page panic %v", recoverVar)
												}
											}()
											// вызываем смену текущего пунтка меню w.tv.SetCurrentItem
											// следом меняем страницу на экране пересоздавая ее из текущей модели редуктора
											if w.Tvm.CurrentMenu != nil {
												if err := w.changePage(); err != nil {
													w.Logger().Errorf("mainwindow tv change page error %s", err.Error())
												}
											}
										},
									},
								},
							},
							dcl.ScrollView{
								StretchFactor: 6,
								MinSize:       dcl.Size{Width: 300},
								Layout: dcl.VBox{MarginsZero: true, SpacingZero: true,
									Margins: dcl.Margins{Left: 0, Top: 0, Right: 5, Bottom: 0}},
								Background: dcl.SolidColorBrush{Color: walk.RGB(255, 255, 255)},
								// StretchFactor: 5,
								Children: []dcl.Widget{
									dcl.Composite{
										AssignTo: &w.pageCom,
										Name:     "placePage",
										Layout:   dcl.VBox{MarginsZero: true, SpacingZero: true},
									},
								},
							},
						},
					},
				},
			},
		},
		StatusBarItems: []dcl.StatusBarItem{
			{
				AssignTo:  &w.SbiLicense,
				Width:     100,
				Text:      "",
				OnClicked: w.SbiLicensePress,
			},
			{
				AssignTo: &w.SbiScan,
				Width:    100,
				// ToolTipText: "Автоматический прием документов в УТМ",
				// OnClicked: w.SbiScanPress,
			},
			{
				AssignTo: &w.SbiFsrarId,
				Width:    120,
				// ToolTipText: "ФСРАР ИД УТМ",
				OnClicked: w.SbiFsrarIdPress,
			},
			{
				AssignTo: &w.SbiUtmState,
				Width:    100,
				// ToolTipText: "Статус УТМ",
				OnClicked: w.SbiUtmPress,
			},
			{
				AssignTo:  &w.SbiState,
				Width:     400,
				Text:      "",
				OnClicked: w.clickHistoryState,
			},
		},
	}).Create(); err != nil {
		return fmt.Errorf(`gui:mainwindow create dcl  %w`, err)
	}
	return nil
}
