package reports

import (
	"fmt"
	"path"

	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"

	"zupper/gui/types"
	"zupper/usecase/services"
)

const modError = "gui:reports"

type ReportsPage struct {
	*walk.Composite
	app    types.IApp
	parent walk.Form
	start  *walk.DateEdit
	end    *walk.DateEdit
}

func NewPage(parent walk.Container, app types.IApp) (pp types.Page, err error) {
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Errorf("%s newHomePage panic %v", modError, r))
		}
	}()
	p := new(ReportsPage)
	p.app = app
	p.parent = parent.Form()

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
						OnDateChanged: func() {
							p.app.SetStartDate(p.start.Date())
						},
					},
					dcl.DateEdit{
						Enabled:  true,
						AssignTo: &p.end,
						// Format: "yyyy-MM-dd",
						Format: "yyyy.MM.dd",
						OnDateChanged: func() {
							p.app.SetEndDate(p.end.Date())
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
									uri := path.Join(p.app.BaseUrl(), "/v1/tranfree")
									p.app.Open(uri)
								},
							},
							dcl.PushButton{
								Text: "Отчет об объемах поставки",
								OnClicked: func() {
									if err := services.New(p.app).HistoryOborotProccess("", ""); err != nil {
										p.app.MessageBox("ошибка", err.Error())
									} else {
										uri := path.Join(p.app.BaseUrl(), "/v1/history/oborot")
										p.app.Open(uri)
									}
								},
							},
							// dcl.PushButton{
							// 	Text: "Интерактивные локальные остатки в браузере",
							// 	OnClicked: func() {
							// 		uri := path.Join(p.app.BaseUrl(), "/v1/localrest")
							// 		p.app.Open(uri)
							// 	},
							// },
							// dcl.PushButton{
							// 	Text: "Диаграмма локальные остатки по типу",
							// 	OnClicked: func() {
							// 		uri := path.Join(p.app.BaseUrl(), "/v1/localrest/chartaptype")
							// 		p.app.Open(uri)
							// 	},
							// },
							dcl.PushButton{
								Text: "История движения АП за период",
								OnClicked: func() {
									uri := path.Join(p.app.BaseUrl(), "/v1/history")
									p.app.Open(uri)
								},
							},
							dcl.PushButton{
								Text: "Реестр списания АП за период",
								OnClicked: func() {
									uri := path.Join(p.app.BaseUrl(), "/v1/history/writeoff")
									p.app.Open(uri)
								},
							},
							dcl.PushButton{
								Text: "Реестр производства АП за период",
								OnClicked: func() {
									uri := path.Join(p.app.BaseUrl(), "/v1/history/production")
									p.app.Open(uri)
								},
							},
							dcl.PushButton{
								Text: "Статистика документов АП (развернуто) за период",
								OnClicked: func() {
									if err := services.New(p.app).HistoryPartsProccess("", "", []string{}); err != nil {
										p.app.MessageBox("ошибка", err.Error())
									} else {
										uri := path.Join(p.app.BaseUrl(), "/v1/history/parts")
										p.app.Open(uri)
									}
								},
							},
							dcl.PushButton{
								Text: "Статистика документов АП (свернуто) за период",
								OnClicked: func() {
									if err := services.New(p.app).HistoryFullProccess("", ""); err != nil {
										p.app.MessageBox("ошибка", err.Error())
									} else {
										uri := path.Join(p.app.BaseUrl(), "/v1/history/full")
										p.app.Open(uri)
									}
								},
							},
							// dcl.PushButton{
							// 	Text: "Отчет об объемах поставки за период",
							// 	OnClicked: func() {
							// 		if err := services.New(p.app).HistoryOborotProccess("", ""); err != nil {
							// 			p.app.MessageBox("ошибка", err.Error())
							// 		} else {
							// 			uri := path.Join(p.app.BaseUrl(), "/v1/history/oborot")
							// 			p.app.Open(uri)
							// 		}
							// 	},
							// },
							// dcl.PushButton{
							// 	Text: "Открыть папку форм",
							// 	OnClicked: func() {
							// 		p.app.OpenDir()
							// 	},
							// },
						}},
					dcl.HSpacer{},
				}},
			// dcl.Composite{
			// 	Layout: dcl.Grid{Columns: 2},
			// 	Children: []dcl.Widget{
			// 		dcl.PushButton{
			// 			Text: "Интерактивные локальные остатки в браузере",
			// 			OnClicked: func() {
			// 				uri := path.Join(p.app.BaseUrl(), "/v1/localrest")
			// 				p.app.Open(uri)
			// 			},
			// 		},
			// 		dcl.HSpacer{Size: 50},
			// 		dcl.PushButton{
			// 			Text: "Диаграмма локальные остатки по типу",
			// 			OnClicked: func() {
			// 				uri := path.Join(p.app.BaseUrl(), "/v1/localrest/chartaptype")
			// 				p.app.Open(uri)
			// 			},
			// 		},
			// 		dcl.HSpacer{Size: 50},
			// 		dcl.PushButton{
			// 			Text: "История движения АП за период",
			// 			OnClicked: func() {
			// 				uri := path.Join(p.app.BaseUrl(), "/v1/history")
			// 				p.app.Open(uri)
			// 			},
			// 		},
			// 		dcl.HSpacer{Size: 50},
			// 		dcl.PushButton{
			// 			Text: "Реестр списания АП за период",
			// 			OnClicked: func() {
			// 				uri := path.Join(p.app.BaseUrl(), "/v1/history/writeoff")
			// 				p.app.Open(uri)
			// 			},
			// 		},
			// 		dcl.HSpacer{Size: 50},
			// 		dcl.PushButton{
			// 			Text: "Статистика документов АП (частями) за период",
			// 			OnClicked: func() {
			// 				if err := services.New(p.app).HistoryPartsProccess("", "", []string{}); err != nil {
			// 					p.app.MessageBox("ошибка", err.Error())
			// 				} else {
			// 					uri := path.Join(p.app.BaseUrl(), "/v1/history/parts")
			// 					p.app.Open(uri)
			// 				}
			// 			},
			// 		},
			// 		dcl.HSpacer{Size: 50},
			// 		dcl.PushButton{
			// 			Text: "Статистика документов АП (один запрос) за период",
			// 			OnClicked: func() {
			// 				if err := services.New(p.app).HistoryFullProccess("", ""); err != nil {
			// 					p.app.MessageBox("ошибка", err.Error())
			// 				} else {
			// 					uri := path.Join(p.app.BaseUrl(), "/v1/history/full")
			// 					p.app.Open(uri)
			// 				}
			// 			},
			// 		},
			// 		dcl.HSpacer{Size: 50},
			// 		dcl.PushButton{
			// 			Text: "Открыть папку форм",
			// 			OnClicked: func() {
			// 				p.app.OpenDir()
			// 			},
			// 		},
			// 	},
			// },
			dcl.VSpacer{},
		},
	}).Create(dcl.NewBuilder(parent)); err != nil {
		return nil, fmt.Errorf("RequestPage %w", err)
	}

	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *ReportsPage) Update() {
}

func (p *ReportsPage) Clear() {
}
