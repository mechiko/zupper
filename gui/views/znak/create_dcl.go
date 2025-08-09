package znak

import (
	"fmt"
	"zupper/domain/models/znakagregate"

	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"
)

func (p *ZnakPage) dclCreate(parent walk.Container, model *znakagregate.ZnakAgregate) error {
	if err := (dcl.Composite{
		Border:   true,
		AssignTo: &p.Composite,
		Name:     "ЧЗ",
		Layout:   dcl.VBox{MarginsZero: false, SpacingZero: false, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
		Children: []dcl.Widget{
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.Label{
						MinSize: dcl.Size{Width: 170},
						Text:    "Штук в группе:",
					},
					dcl.ComboBox{
						AssignTo:              &p.ipsCombo,
						Alignment:             dcl.AlignHNearVNear,
						Editable:              false,
						Value:                 "",
						Model:                 itemPerPage,
						OnCurrentIndexChanged: func() {},
					},
					dcl.Label{
						AssignTo: &p.waitStateLbl,
						Text:     "идет расчет аггрегации тары ...",
					},
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.PushButton{
						MinSize: dcl.Size{Width: 170},
						Text:    "Заказ Коробка",
						// OnClicked: p.selectGroupDialog,
					},
					dcl.Label{
						AssignTo: &p.groupLbl,
					},
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.Label{
						AssignTo: &p.groupItogLbl,
						Text:     "",
					},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.PushButton{
						MinSize: dcl.Size{Width: 170},
						Text:    "Заказ Бутылка",
						// OnClicked: p.selectPackageDialog,
					},
					dcl.Label{
						AssignTo: &p.packageLbl,
					},
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.Label{
						AssignTo: &p.packageItogLbl,
						Text:     "",
					},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.PushButton{
						AssignTo:  &p.filePb,
						MinSize:   dcl.Size{Width: 170},
						Enabled:   false,
						Text:      "Выгрузить файл (китай)",
						OnClicked: func() {},
					},
					dcl.Label{
						AssignTo: &p.fileLbl,
						// TODO: Add specific label field for this button
					},
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.PushButton{
						AssignTo:  &p.filePbXml,
						MinSize:   dcl.Size{Width: 170},
						Enabled:   false,
						Text:      "Выгрузить файл XML агрегации (ЧЗ)",
						OnClicked: func() {},
					},
					dcl.Label{
						AssignTo: &p.fileLbl,
						// TODO: Add specific label field for this button
					},
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.PushButton{
						AssignTo:  &p.filePbA3,
						MinSize:   dcl.Size{Width: 170},
						Enabled:   false,
						Text:      "Выгрузить файл для A3 агрегации (ексель набор)",
						OnClicked: func() {},
					},
					dcl.Label{
						AssignTo: &p.fileLbl,
						// TODO: Add specific label field for this button
					},
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.PushButton{
						AssignTo:  &p.filePb1C,
						MinSize:   dcl.Size{Width: 170},
						Enabled:   false,
						Text:      "Выгрузить файл для 1с агрегации (ексель)",
						OnClicked: func() {},
					},
					dcl.Label{
						AssignTo: &p.fileLbl,
						// TODO: Add specific label field for this button
					},
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.PushButton{
						AssignTo:  &p.filePbCsv,
						MinSize:   dcl.Size{Width: 170},
						Enabled:   false,
						Text:      "Выгрузить файл для 1с агрегации (CSV)",
						OnClicked: func() {},
					},
					dcl.Label{
						AssignTo: &p.fileLbl,
						// TODO: Add specific label field for this button
					},
					dcl.HSpacer{},
				},
			},
		},
	}).Create(dcl.NewBuilder(parent)); err != nil {
		return fmt.Errorf("gui:view znak %w", err)
	}

	return nil
}
