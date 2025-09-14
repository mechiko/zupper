package setup

import (
	"fmt"
	"zupper/domain/models/application"
	"zupper/reductor"

	"github.com/mechiko/utility"

	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"
)

func (p *SetupPage) dclCreate(parent walk.Container, model *application.Application) error {
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
						AssignTo:  &p.browserCB,
						Alignment: dcl.AlignHNearVCenter,
						Name:      "combobox",
						// Background: dcl.SolidColorBrush{},
						// Background:            dcl.SolidColorBrush{Color: walk.RGB(0x34, 0x82, 0xeb)},
						Editable: false,
						Value:    model.Browser,
						Model:    []string{string(utility.Default), string(utility.Chrome), string(utility.Firefox), string(utility.Yandex), string(utility.Edge)},
						OnCurrentIndexChanged: func() {
							p.Logger().Debug("browser current index change")
							txt := p.browserCB.Text()
							modelChange, _ := p.Model()
							modelChange.Browser = utility.Browser(txt)
							err := modelChange.SyncToStore(p)
							if err != nil {
								p.Logger().Errorf("change browser set in store error %s", err.Error())
							}
							err = modelChange.Save(p)
							if err != nil {
								p.Logger().Errorf("change browser set in store error %s", err.Error())
							}
							err = reductor.Instance().SetModel(modelChange, false)
							if err != nil {
								p.Logger().Errorf("change browser set in reductor error %s", err.Error())
							}
						},
					},
					// 		dcl.HSpacer{Size: 20},
					// 		dcl.PushButton{
					// 			Text: "Открыть Веб Приложение",
					// 			OnClicked: func() {
					// 				// uri := path.Join(p.BaseUrl(), "/v1/home")
					// 				// p.Open(uri)
					// 			},
					// 		},
					// 		dcl.HSpacer{},
				},
			},
			// dcl.GroupBox{
			// 	Title:  "УТМ",
			// 	Layout: dcl.VBox{MarginsZero: false, SpacingZero: false, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
			// 	Children: []dcl.Widget{
			// 		dcl.Composite{
			// 			Layout:    dcl.HBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 0, Top: 0, Right: 0, Bottom: 0}},
			// 			Border:    false,
			// 			Alignment: dcl.Alignment2D(walk.AlignHNearVNear),
			// 			Children: []dcl.Widget{
			// 				dcl.Label{
			// 					Text: "Хост:",
			// 				},
			// 				dcl.LineEdit{
			// 					AssignTo: &p.utmhost,
			// 					MaxSize:  dcl.Size{Width: 200},
			// 					Text:     model.Host,
			// 				},
			// 				dcl.HSpacer{},
			// 			},
			// 		},
			// 		dcl.Composite{
			// 			Layout:    dcl.HBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 0, Top: 0, Right: 0, Bottom: 0}},
			// 			Border:    false,
			// 			Alignment: dcl.Alignment2D(walk.AlignHNearVNear),
			// 			Children: []dcl.Widget{
			// 				dcl.Label{
			// 					Text: "Порт:",
			// 				},
			// 				dcl.LineEdit{
			// 					AssignTo: &p.utmport,
			// 					MaxSize:  dcl.Size{Width: 200},
			// 					Text:     model.Port,
			// 				},
			// 				dcl.HSpacer{},
			// 			},
			// 		},
			// 	}},
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
								Text: "БД программы:",
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
					// dcl.PushButton{
					// 	AssignTo: &p.saveconf,
					// 	// ColumnSpan: 2,
					// 	Text: "Обновить конфигурацию",
					// 	// OnClicked: p.saveConfig,
					// 	OnClicked: func() {
					// 		// отправляем обновление своей модели в канал
					// 		if p.sendChan != nil {
					// 			p.sendChan(p.model)
					// 		}
					// 	},
					// },
					dcl.PushButton{
						Text: "Открыть папку выгрузки",
						OnClicked: func() {
							if err := utility.OpenFileInShell(p.Options().Output); err != nil {
								utility.MessageBox("ошибка", err.Error())
							}
						},
					},
					dcl.PushButton{
						Text: "Открыть папку настройки и логов",
						OnClicked: func() {
							if err := utility.OpenFileInShell(p.ConfigPath()); err != nil {
								utility.MessageBox("ошибка", err.Error())
							}
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
