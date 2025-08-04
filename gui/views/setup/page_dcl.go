package setup

import (
	"fmt"

	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"
)

func (p *HomePage) dclCreate(parent walk.Container, model *SetupModel) error {
	if err := (dcl.Composite{
		AssignTo: &p.Composite,
		Name:     model.Title,
		// Background: dcl.SolidColorBrush{walk.RGB(255, 255, 255)},
		Layout:    dcl.VBox{Spacing: 10, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
		Alignment: dcl.Alignment2D(walk.AlignHNearVNear),
		Border:    true,
		Children: []dcl.Widget{
			dcl.Composite{
				Border: false,
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 0}},
				Children: []dcl.Widget{
					dcl.Label{
						Text: "Выберите браузер:",
						// Background: dcl.SolidColorBrush{Color: walk.RGB(0x34, 0x82, 0xeb)},
						// TextAlignment: dcl.AlignNear,
					},
					dcl.ComboBox{
						Alignment: dcl.AlignHNearVCenter,
						// Background: dcl.SolidColorBrush{},
						// Background:            dcl.SolidColorBrush{Color: walk.RGB(0x34, 0x82, 0xeb)},
						Editable:              false,
						Value:                 model.Browser,
						Model:                 []string{"", "Chrome", "Firefox", "Yandex", "MSEdge"},
						OnCurrentIndexChanged: p.changeIndexBrowser,
					},
					dcl.HSpacer{Size: 20},
					dcl.PushButton{
						Text: "Открыть Веб Приложение",
						OnClicked: func() {
							// uri := path.Join(p.BaseUrl(), "/v1/home")
							// p.Open(uri)
						},
					},
					dcl.HSpacer{},
				},
			},
			dcl.GroupBox{
				Title:  "УТМ",
				Layout: dcl.VBox{MarginsZero: false, SpacingZero: false, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.Composite{
						Layout:    dcl.HBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 0, Top: 0, Right: 0, Bottom: 0}},
						Border:    false,
						Alignment: dcl.Alignment2D(walk.AlignHNearVNear),
						Children: []dcl.Widget{
							dcl.Label{
								Text: "Хост:",
							},
							dcl.LineEdit{
								AssignTo: &p.utmhost,
								MaxSize:  dcl.Size{Width: 200},
								Text:     model.Host,
							},
							dcl.HSpacer{},
						},
					},
					dcl.Composite{
						Layout:    dcl.HBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 0, Top: 0, Right: 0, Bottom: 0}},
						Border:    false,
						Alignment: dcl.Alignment2D(walk.AlignHNearVNear),
						Children: []dcl.Widget{
							dcl.Label{
								Text: "Порт:",
							},
							dcl.LineEdit{
								AssignTo: &p.utmport,
								MaxSize:  dcl.Size{Width: 200},
								Text:     model.Port,
							},
							dcl.HSpacer{},
						},
					},
				}},
			dcl.GroupBox{
				Title:  "Конфигураци БД",
				Layout: dcl.VBox{MarginsZero: false, SpacingZero: false, Margins: dcl.Margins{Left: 10, Top: 10, Right: 10, Bottom: 10}},
				Children: []dcl.Widget{
					dcl.Composite{
						Layout:    dcl.HBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 0, Top: 0, Right: 0, Bottom: 0}},
						Border:    false,
						Alignment: dcl.Alignment2D(walk.AlignHNearVNear),
						Children: []dcl.Widget{
							dcl.Label{
								Text: "БД Config:",
							},
							dcl.Label{
								AssignTo: &p.lblDbConfig,
								Text:     model.DbConfigDesc,
							},
							dcl.HSpacer{},
						},
					},
					dcl.Composite{
						Layout:    dcl.HBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 0, Top: 0, Right: 0, Bottom: 0}},
						Border:    false,
						Alignment: dcl.Alignment2D(walk.AlignHNearVNear),
						Children: []dcl.Widget{
							dcl.Label{
								Text: "БД A3:",
							},
							dcl.Label{
								AssignTo: &p.lblDbA3,
								Text:     model.DbA3Desc,
							},
							dcl.HSpacer{},
						},
					},
					dcl.Composite{
						Layout:    dcl.HBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 0, Top: 0, Right: 0, Bottom: 0}},
						Border:    false,
						Alignment: dcl.Alignment2D(walk.AlignHNearVNear),
						Children: []dcl.Widget{
							dcl.Label{
								Text: "БД ЧЗ:",
							},
							dcl.Label{
								AssignTo: &p.lblDbZnak,
								Text:     model.DbZnakDesc,
							},
							dcl.HSpacer{},
						},
					},
					dcl.Composite{
						Layout:    dcl.HBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 0, Top: 0, Right: 0, Bottom: 0}},
						Border:    false,
						Alignment: dcl.Alignment2D(walk.AlignHNearVNear),
						Children: []dcl.Widget{
							dcl.Label{
								Text: "БД AlcoGo4Lite:",
							},
							dcl.Label{
								AssignTo: &p.lblDbLite,
								Text:     model.DbLiteDesc,
							},
							dcl.HSpacer{},
						},
					},
				}},
			dcl.Composite{
				Border: false,
				Layout: dcl.VBox{MarginsZero: true, SpacingZero: true, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 0}},
				Children: []dcl.Widget{
					dcl.PushButton{
						AssignTo: &p.saveconf,
						// ColumnSpan: 2,
						Text: "Обновить конфигурацию",
						// OnClicked: p.saveConfig,
						OnClicked: func() {
							// отправляем обновление своей модели в канал
							p.sendChan(p.model)
						},
					},
					dcl.PushButton{
						Text: "Открыть папку выгрузки",
						OnClicked: func() {
							// p.OpenDir()
						},
					},
				}},
			dcl.VSpacer{},
		},
	}).Create(dcl.NewBuilder(parent)); err != nil {
		return fmt.Errorf("view:setup dcl create %w", err)
	}
	return nil
}
