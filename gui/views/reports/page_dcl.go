package reports

import (
	"fmt"
	"zupper/domain/models/application"
	"zupper/reductor"

	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"
)

func (p *ReportsPage) dclCreate(parent walk.Container, model *application.Application) error {
	if err := (dcl.Composite{
		AssignTo:  &p.Composite,
		Border:    true,
		Layout:    dcl.VBox{MarginsZero: true, SpacingZero: true, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
		MaxSize:   dcl.Size{Width: 400},
		MinSize:   dcl.Size{Width: 400},
		Alignment: dcl.Alignment2D(walk.AlignHNearVNear),
		Children: []dcl.Widget{
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.Label{
						Text: "Периода расчета:",
					},
					dcl.DateEdit{
						Enabled:  true,
						AssignTo: &p.start,
						Format:   "yyyy.MM.dd",
						Date:     model.StartDate(),
						OnDateChanged: func() {
							modelReductor, _ := reductor.Instance().Model(model.Model())
							modelTmp, ok := modelReductor.(*application.Application)
							if !ok {
								return
							}
							modelTmp.SetStartDate(p.start.Date())
							reductor.Instance().SetModel(modelTmp, false)
						},
					},
					dcl.DateEdit{
						Enabled:  true,
						AssignTo: &p.end,
						// Format: "yyyy-MM-dd",
						Format: "yyyy.MM.dd",
						Date:   model.EndDate(),
						OnDateChanged: func() {
							modelReductor, _ := reductor.Instance().Model(model.Model())
							modelTmp, ok := modelReductor.(*application.Application)
							if !ok {
								return
							}
							modelTmp.SetEndDate(p.end.Date())
							reductor.Instance().SetModel(modelTmp, false)
						},
					},
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				// Alignment: dcl.AlignHNearVCenter,
				Children: []dcl.Widget{
					dcl.HSpacer{},
				}},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				// MinSize: dcl.Size{Width: 600},
				Children: []dcl.Widget{
					dcl.Composite{
						Layout:  dcl.VBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
						MinSize: dcl.Size{Width: 600},
						Children: []dcl.Widget{
							dcl.PushButton{
								Text: "Отчет об объемах по видам (пиво) строгий",
								OnClicked: func() {
									// uri := path.Join(p.app.BaseUrl(), "/v1/tranfree")
									// p.app.Open(uri)
								},
							},
							dcl.PushButton{
								Text: "Отчет об объемах поставки",
								OnClicked: func() {
									// if err := services.New(p.app).HistoryOborotProccess("", ""); err != nil {
									// 	p.app.MessageBox("ошибка", err.Error())
									// } else {
									// 	uri := path.Join(p.app.BaseUrl(), "/v1/history/oborot")
									// 	p.app.Open(uri)
									// }
								},
							},
							dcl.PushButton{
								Text: "История движения АП за период",
								OnClicked: func() {
									// uri := path.Join(p.app.BaseUrl(), "/v1/history")
									// p.app.Open(uri)
								},
							},
							dcl.PushButton{
								Text: "Реестр списания АП за период",
								OnClicked: func() {
									// uri := path.Join(p.app.BaseUrl(), "/v1/history/writeoff")
									// p.app.Open(uri)
								},
							},
							dcl.PushButton{
								Text: "Реестр производства АП за период",
								OnClicked: func() {
									// uri := path.Join(p.app.BaseUrl(), "/v1/history/production")
									// p.app.Open(uri)
								},
							},
							dcl.PushButton{
								Text: "Статистика документов АП (развернуто) за период",
								OnClicked: func() {
									// if err := services.New(p.app).HistoryPartsProccess("", "", []string{}); err != nil {
									// 	p.app.MessageBox("ошибка", err.Error())
									// } else {
									// 	uri := path.Join(p.app.BaseUrl(), "/v1/history/parts")
									// 	p.app.Open(uri)
									// }
								},
							},
							dcl.PushButton{
								Text: "Статистика документов АП (свернуто) за период",
								OnClicked: func() {
									// if err := services.New(p.app).HistoryFullProccess("", ""); err != nil {
									// 	p.app.MessageBox("ошибка", err.Error())
									// } else {
									// 	uri := path.Join(p.app.BaseUrl(), "/v1/history/full")
									// 	p.app.Open(uri)
									// }
								},
							},
						}},
					dcl.HSpacer{},
				}},
			dcl.VSpacer{},
		},
	}).Create(dcl.NewBuilder(parent)); err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	return nil
}
