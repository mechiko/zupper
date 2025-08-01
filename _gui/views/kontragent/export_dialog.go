package kontragent

import (
	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"
)

func (p *KontragentListPage) ExportDialog() (int, error) {
	var dlgExport *walk.Dialog
	var dlgExportClosePb *walk.PushButton
	var dlgExportOpenDirPb *walk.PushButton
	var state *walk.TableView

	modelState := NewListStateModel(p.app)
	// stateInfo := make(StateItemSlice, 0)
	// info := &StateItem{
	// 	Title: "экспорт данных по партнеру",
	// 	Info:  "",
	// 	Color: "black",
	// }
	// stateInfo = append(stateInfo, info)
	// modelState.SetItems(stateInfo)
	if err := (dcl.Dialog{
		AssignTo: &dlgExport,
		Title:    "Экспорт",
		MinSize:  dcl.Size{Width: 200, Height: 300},
		// Layout:  VBox{MarginsZero: true, SpacingZero: true, Alignment: Alignment2D(AlignHNearVNear)},
		Layout: dcl.VBox{
			MarginsZero: true,
			SpacingZero: true,
			Alignment:   dcl.Alignment2D(dcl.AlignHNearVNear),
		},
		Children: []dcl.Widget{
			dcl.Composite{
				Layout: dcl.Grid{
					Columns:   1,
					Alignment: dcl.Alignment2D(dcl.AlignHNearVNear),
				},
				Children: []dcl.Widget{
					dcl.Label{
						Text:      "Экспорт данных для контрагента",
						Alignment: dcl.Alignment2D(dcl.AlignCenter),
					},
					dcl.TableView{
						AssignTo:                 &state,
						Name:                     "state",
						AlternatingRowBG:         true,
						CheckBoxes:               false,
						ColumnsOrderable:         false,
						NotSortableByHeaderClick: true,
						MultiSelection:           false,
						LastColumnStretched:      true,
						HeaderHidden:             true,
						Columns: []dcl.TableViewColumn{
							{Title: "Состояние", Alignment: dcl.AlignNear},
						},
						StyleCell: func(style *walk.CellStyle) {
							item := modelState.items[style.Row()]
							switch item.Color {
							case "red":
								style.TextColor = red
							default:
								style.TextColor = black
							}
						},
						Model: modelState,
						OnItemActivated: func() {
						},
					},
				},
			},
			dcl.Composite{
				Layout:    dcl.HBox{},
				Alignment: dcl.AlignHFarVFar,
				// Alignment:     AlignHNearVNear,
				StretchFactor: 1,
				Children: []dcl.Widget{
					dcl.HSpacer{},
					dcl.PushButton{
						AssignTo: &dlgExportOpenDirPb,
						Text:     "Открыть папку выгрузки",
						OnClicked: func() {
							p.app.OpenDir()
						},
						Enabled: false,
						Visible: true,
					},
					dcl.PushButton{
						AssignTo:  &dlgExportClosePb,
						Text:      "Закрыть",
						OnClicked: func() { dlgExport.Accept() },
						Enabled:   false,
						Visible:   true,
					},
				},
			},
		},
	}).Create(p.parent); err != nil {
		return 0, err
	}
	dlgExport.Starting().Attach(func() {
		go func() {
			k := p.model.filtered[p.tv.CurrentIndex()]
			if err := p.ExportExcelOborot(k); err != nil {
				p.app.MessageBox("ошибка", err.Error())
				modelState.AddState("ошибка экспорта файла excel", "red")
			} else {
				modelState.AddState("экспорт файла excel завершен", "")
			}
			p.Synchronize(func() {
				modelState.ResetRows()
			})
			if err := p.ExportFirmShipper(k); err != nil {
				p.app.MessageBox("ошибка", err.Error())
				modelState.AddState("ошибка экспорта файла xml", "red")
			} else {
				modelState.AddState("экспорт файла xml организаций завершен", "")
			}
			p.Synchronize(func() {
				modelState.ResetRows()
			})
			if err := p.ExporеPartnerOborot(k); err != nil {
				p.app.MessageBox("ошибка", err.Error())
				modelState.AddState("ошибка экспорта файла xml", "red")
			} else {
				modelState.AddState("экспорт файла xml оборотов завершен", "")
			}
			p.Synchronize(func() {
				modelState.ResetRows()
			})
			dlgExportClosePb.SetEnabled(true)
			dlgExportOpenDirPb.SetEnabled(true)
			// p.Update()
		}()
	})
	ret := dlgExport.Run()

	dlgExport.Dispose()
	return ret, nil
}
