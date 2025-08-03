package mainwindow

import (
	"fmt"
	"image/color"

	"zupper/gui/resource"

	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"
)

func (w *MainWindow) Create() error {
	var err error
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

	if err := (dcl.MainWindow{
		AssignTo: &w.MainWindow,
		Name:     w.Cfg.Name,
		Title:    w.Cfg.Title,
		Enabled:  w.Cfg.Enabled,
		// Visible:  false,
		Font:    w.Cfg.Font,
		MinSize: w.Cfg.MinSize,
		// MaxSize:          w.Cfg.MaxSize,
		// Size:             dcl.Size{Width: 1000, Height: 700},
		Size:             w.Cfg.MinSize,
		MenuItems:        w.Cfg.MenuItems,
		ToolBar:          w.Cfg.ToolBar,
		ContextMenuItems: w.Cfg.ContextMenuItems,
		OnKeyDown:        w.Cfg.OnKeyDown,
		OnKeyPress:       w.Cfg.OnKeyPress,
		OnKeyUp:          w.Cfg.OnKeyUp,
		OnMouseDown:      w.Cfg.OnMouseDown,
		OnMouseMove:      w.Cfg.OnMouseMove,
		OnMouseUp:        w.Cfg.OnMouseUp,
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
												if r := recover(); r != nil {
													w.Logger().Errorf("TV change page panic %v", r)
													panic(fmt.Errorf("TV change page panic %v", r))
													// TODO Shutdown
												}
											}()
											if err := w.сhangePage(); err != nil {
												w.Logger().Errorf("TV change page error %s", err.Error())
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
		return fmt.Errorf(`w.Create() %w`, err)
	}
	succeeded := false
	defer func() {
		if !succeeded {
			w.Dispose()
		}
	}()

	w.tv.SetCurrentItem(w.Tvm.DefaultMenu())
	w.CurrentPageChanged().Attach(w.Cfg.OnCurrentPageChanged)
	succeeded = true
	w.сhangePage()
	return err
}
