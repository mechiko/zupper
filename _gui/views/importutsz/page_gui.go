package importutsz

import (
	"fmt"
	_ "time/tzdata"

	"zupper/gui/types"

	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"
)

func NewPage(parent walk.Container, app types.IApp) (types.Page, error) {
	// func NewPage(parent walk.Container, a entity.IApp, close func()) (c *walk.Composite, err error) {
	p := new(ImportTTNPage)
	p.form = parent.Form()
	p.app = app

	if err := p.init(); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	if err := (dcl.Composite{
		AssignTo:  &p.Composite,
		Border:    true,
		MinSize:   dcl.Size{Width: 400, Height: 300},
		Layout:    dcl.VBox{MarginsZero: false, SpacingZero: false, Margins: dcl.Margins{Left: 0, Top: 0, Right: 5, Bottom: 0}},
		Alignment: dcl.Alignment2D(walk.AlignHNearVNear),
		Children: []dcl.Widget{
			dcl.Composite{
				Border:    false,
				Layout:    dcl.VBox{MarginsZero: false, SpacingZero: false, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Alignment: dcl.AlignHNearVNear,
				Children: []dcl.Widget{
					dcl.Composite{
						DoubleBuffering: true,
						Layout:          dcl.Grid{Rows: 1, MarginsZero: true},
						MaxSize:         dcl.Size{Width: 700},
						Children: []dcl.Widget{
							dcl.PushButton{
								AssignTo:  &p.btnSrc,
								Text:      "Выберите файл",
								MaxSize:   dcl.Size{Width: 150},
								OnClicked: p.openSrcDlg,
							},
							dcl.Label{
								AssignTo:      &p.lblFile,
								TextAlignment: dcl.AlignNear,
							},
							dcl.HSpacer{},
						},
					},
					dcl.Composite{
						DoubleBuffering: true,
						Layout:          dcl.Grid{Rows: 1, MarginsZero: true},
						Children: []dcl.Widget{
							dcl.Label{
								Alignment: dcl.AlignHNearVCenter,
								Text:      "Кол-во ТТН:",
							},
							dcl.Label{
								AssignTo:  &p.lblCountTtn,
								Alignment: dcl.AlignHNearVCenter,
							},
							dcl.Label{
								AssignTo:  &p.lblProccessingError,
								Alignment: dcl.AlignHNearVCenter,
							},
							dcl.HSpacer{},
						},
					},
					// dcl.Composite{
					// 	Layout: dcl.Grid{Rows: 1, MarginsZero: true, Margins: dcl.Margins{Left: 0, Right: 5, Bottom: 0, Top: 0}},
					// 	Children: []dcl.Widget{
					// 		dcl.Label{
					// 			Text:          "фильтр остатков по Дате розлива (YYYY.MM.DD) (только по этому списку):",
					// 			TextAlignment: dcl.AlignNear,
					// 		},
					// 		dcl.LineEdit{
					// 			AssignTo:  &p.filterBottlingFld,
					// 			Alignment: dcl.AlignHNearVNear,
					// 			MaxSize:   dcl.Size{Width: 400},
					// 			Text:      "",
					// 			OnTextChanged: func() {
					// 				p.app.ImportTTN().SetFilterBolling(p.filterBottlingFld.Text())
					// 			},
					// 		},
					// 		dcl.HSpacer{Size: 10},
					// 	},
					// },
					// dcl.Composite{
					// 	Layout: dcl.Grid{Rows: 1, MarginsZero: true, Margins: dcl.Margins{Left: 0, Right: 5, Bottom: 0, Top: 0}},
					// 	Children: []dcl.Widget{
					// 		dcl.Label{
					// 			Text:          "используем остатки, начиная с Даты розлива (YYYY.MM.DD):",
					// 			TextAlignment: dcl.AlignNear,
					// 		},
					// 		dcl.LineEdit{
					// 			AssignTo:  &p.startBottlingFld,
					// 			Alignment: dcl.AlignHNearVNear,
					// 			MaxSize:   dcl.Size{Width: 400},
					// 			Text:      "",
					// 			OnTextChanged: func() {
					// 				p.app.ImportTTN().SetStartBotlling(p.startBottlingFld.Text())
					// 			},
					// 		},
					// 		dcl.HSpacer{Size: 10},
					// 	},
					// },
					dcl.Composite{
						DoubleBuffering: true,
						Layout:          dcl.Grid{Rows: 1, MarginsZero: true},
						Children: []dcl.Widget{
							dcl.CheckBox{
								AssignTo:  &p.fifoCheckBox,
								Alignment: dcl.AlignHNearVNear,
								Text:      "очередь списания остатков (сортировка Дата розлива -> Количество)",
								OnCheckedChanged: func() {
									p.app.ImportTTN().SetFifo(p.fifoCheckBox.CheckState() == 1)
								},
							},
							dcl.HSpacer{},
						},
					},
					dcl.Composite{
						DoubleBuffering: true,
						Layout:          dcl.Grid{Rows: 1, MarginsZero: true},
						Children: []dcl.Widget{
							dcl.CheckBox{
								AssignTo:  &p.splitCheckBox,
								Alignment: dcl.AlignHNearVCenter,
								Text:      "возможно подбирать из разных партий остатков",
								OnCheckedChanged: func() {
									p.app.ImportTTN().SetSplit(p.splitCheckBox.CheckState() == 1)
								},
							},
							dcl.HSpacer{},
						},
					},
					dcl.Composite{
						DoubleBuffering: true,
						Layout:          dcl.Grid{Rows: 1, MarginsZero: true},
						Children: []dcl.Widget{
							dcl.CheckBox{
								AssignTo:  &p.reimportCheckBox,
								Alignment: dcl.AlignHNearVCenter,
								Text:      "при повторном импорте черновики импортированных ТТН удаляются",
								OnCheckedChanged: func() {
									p.app.ImportTTN().SetReImport(p.reimportCheckBox.CheckState() == 1)
								},
							},
							dcl.HSpacer{},
						},
					},
					// dcl.Composite{
					// 	DoubleBuffering: true,
					// 	Layout:          dcl.Grid{Rows: 1, MarginsZero: true},
					// 	Children: []dcl.Widget{
					// 		dcl.CheckBox{
					// 			AssignTo:  &p.ignoreRestCheckBox,
					// 			Alignment: dcl.AlignHNearVCenter,
					// 			Text:      "игнорировать наличие количества по коду АП в остатках",
					// 			OnCheckedChanged: func() {
					// 				p.app.ImportTTN().SetIgnoreRest(p.ignoreRestCheckBox.CheckState() == 1)
					// 			},
					// 		},
					// 		dcl.HSpacer{},
					// 	},
					// },
					dcl.Composite{
						DoubleBuffering: true,
						Layout:          dcl.Grid{Rows: 1, MarginsZero: true},
						Children: []dcl.Widget{
							dcl.CheckBox{
								AssignTo:  &p.emptyTtnCheckBox,
								Alignment: dcl.AlignHNearVCenter,
								Text:      "импортировать пустые ТТН",
								OnCheckedChanged: func() {
									p.app.ImportTTN().SetEmptyTtn(p.emptyTtnCheckBox.CheckState() == 1)
								},
							},
							dcl.HSpacer{},
						},
					},
					dcl.Composite{
						Layout: dcl.Grid{Rows: 1, Margins: dcl.Margins{Left: 0, Top: 0, Right: 0, Bottom: 10}},
						Children: []dcl.Widget{
							dcl.PushButton{
								AssignTo:  &p.btnExamen,
								MaxSize:   dcl.Size{Width: 120},
								Text:      "Анализ",
								Enabled:   false,
								OnClicked: p.clickExamen,
							},
							dcl.PushButton{
								AssignTo:  &p.btnImportTtn,
								MaxSize:   dcl.Size{Width: 120},
								Text:      "Импорт",
								Enabled:   false,
								OnClicked: p.clickImportTtn,
							},
							dcl.PushButton{
								AssignTo:  &p.btnControlTtn,
								MaxSize:   dcl.Size{Width: 120},
								Enabled:   false,
								Text:      "Контроль",
								OnClicked: p.openImportCheck,
							},
							// dcl.PushButton{
							// 	AssignTo: &p.btnCancel,
							// 	MaxSize:  dcl.Size{Width: 100},
							// 	Enabled:  true,
							// 	Text:     "Выход",
							// 	// OnClicked: func() {
							// 	// 	p.app.Shutdown()
							// 	// },
							// },
						},
					},
					dcl.VSpacer{},
				},
			},
		},
	}).Create(dcl.NewBuilder(parent)); err != nil {
		return nil, fmt.Errorf("Home_Page.Create()%w", err)
	}
	return p, nil
}

func (p *ImportTTNPage) Clear() {
}
