package maintain

import (
	"fmt"
	"path"

	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"

	"zupper/gui/types"
)

const modError = "gui:maintain"

type MaintainPage struct {
	*walk.Composite
	app    types.IApp
	parent walk.Form
	start  *walk.DateEdit
	end    *walk.DateEdit
}

func NewPage(parent walk.Container, app types.IApp) (pp types.Page, err error) {
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Errorf("%s newPage panic %v", modError, r))
		}
	}()
	p := new(MaintainPage)
	p.app = app
	p.parent = parent.Form()

	if err := (dcl.Composite{
		AssignTo:  &p.Composite,
		Layout:    dcl.VBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
		Border:    true,
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
			dcl.GroupBox{
				Title:     "БД АлкоХелп 3",
				Layout:    dcl.VBox{Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Alignment: dcl.AlignHNearVCenter,
				Children: []dcl.Widget{
					dcl.Composite{
						Layout: dcl.HBox{Spacing: 40, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
						Children: []dcl.Widget{
							dcl.PushButton{
								MaxSize: dcl.Size{Width: 250},
								Text:    "Проверка БД",
								OnClicked: func() {
									uri := path.Join(p.app.BaseUrl(), "/v1/maintain/adminreport")
									p.app.Open(uri)
								},
							},
							dcl.TextLabel{
								MaxSize: dcl.Size{Width: 600},
								Text:    "проверка базы данных алкохелп 3 на предмет дублирования данных, что может приводить к задвоению объемов",
							},
							dcl.HSpacer{},
						},
					},
					dcl.Composite{
						Layout: dcl.HBox{Spacing: 40, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
						Children: []dcl.Widget{
							dcl.PushButton{
								MaxSize: dcl.Size{Width: 250},
								Text:    "Список проблемных ТТН",
								OnClicked: func() {
									uri := path.Join(p.app.BaseUrl(), "/v1/defect/ttn")
									p.app.Open(uri)
								},
							},
							dcl.TextLabel{
								MaxSize: dcl.Size{Width: 600},
								Text:    "вывод списка проблемных (отказы, даты, расхождения) ТТН (отгрузка) за период",
							},
							dcl.HSpacer{},
						},
					},
					dcl.Composite{
						Layout: dcl.HBox{Spacing: 40, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
						Children: []dcl.Widget{
							dcl.PushButton{
								MaxSize: dcl.Size{Width: 250},
								Text:    "Удалить запросы",
								OnClicked: func() {
									if result, err := p.app.Repo().DbA3().AdminRequestRemove(); err != nil {
										p.app.Logger().Errorf("%s %s", modError, err.Error())
										p.app.MessageBox("ошибка БД", err.Error())
									} else {
										if result != 0 {
											p.app.MessageBox("Сообщение", fmt.Sprintf("Записи удалены %d", result))
										} else {
											p.app.MessageBox("Сообщение", "Записи удалены")
										}
									}
								},
							},
							dcl.TextLabel{
								MaxSize: dcl.Size{Width: 600},
								Text:    "удаление запросов в статусе Отправлен и Ошибка, они возможно блокируют получение справок по движение АП",
							},
							dcl.HSpacer{},
						},
					},
					// очистка порционных остатков в таком виде опасна, не ходят потом ТТН ... акты
					// dcl.Composite{
					// 	Layout: dcl.HBox{Spacing: 40, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
					// 	Children: []dcl.Widget{
					// 		dcl.PushButton{
					// 			MaxSize: dcl.Size{Width: 250},
					// 			Text:    "Очистка порционных остатков. Необходимо запросить остатки и выполнить синхронизацию!",
					// 			OnClicked: func() {
					// 				if err := p.app.Repo().DbA3().AdminClearRestVolume(); err != nil {
					// 					p.app.MessageBox("Ошибка БД А3", err.Error())
					// 				} else {
					// 					p.app.MessageBox("Success", "Порционные остатки очищены")
					// 				}
					// 			},
					// 		},
					// 		dcl.TextLabel{
					// 			MaxSize: dcl.Size{Width: 600},
					// 			Text:    "очистка остатков связанных с порционным списанием, полностью удаляются",
					// 		},
					// 		dcl.HSpacer{},
					// 	},
					// },
					dcl.Composite{
						Layout: dcl.HBox{Spacing: 40, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
						Children: []dcl.Widget{
							dcl.PushButton{
								Text: "Интерактивные локальные остатки в браузере",
								OnClicked: func() {
									uri := path.Join(p.app.BaseUrl(), "/v1/localrest")
									p.app.Open(uri)
								},
							},
							dcl.TextLabel{
								MaxSize: dcl.Size{Width: 600},
								Text:    "",
							},
							dcl.HSpacer{},
						},
					},
					dcl.Composite{
						Layout: dcl.HBox{Spacing: 40, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
						Children: []dcl.Widget{
							dcl.PushButton{
								Text: "Диаграмма локальные остатки по типу",
								OnClicked: func() {
									uri := path.Join(p.app.BaseUrl(), "/v1/localrest/chartaptype")
									p.app.Open(uri)
								},
							},
							dcl.TextLabel{
								MaxSize: dcl.Size{Width: 600},
								Text:    "",
							},
							dcl.HSpacer{},
						},
					},
					// dcl.Composite{
					// 	Layout: dcl.HBox{Spacing: 40, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
					// 	Children: []dcl.Widget{
					// 		dcl.PushButton{
					// 			Text: "История движения АП за период",
					// 			OnClicked: func() {
					// 				uri := path.Join(p.app.BaseUrl(), "/v1/history")
					// 				p.app.Open(uri)
					// 			},
					// 		},
					// 		dcl.TextLabel{
					// 			MaxSize: dcl.Size{Width: 600},
					// 			Text:    "",
					// 		},
					// 		dcl.HSpacer{},
					// 	},
					// },
				},
			},
			// dcl.Composite{
			// 	AssignTo:  &p.Composite,
			// 	Layout:    dcl.HBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 0, Top: 0, Right: 0, Bottom: 0}},
			// 	Border:    false,
			// 	Alignment: dcl.Alignment2D(walk.AlignHNearVNear),
			// 	Children: []dcl.Widget{
			// 		dcl.PushButton{
			// 			MaxSize: dcl.Size{Width: 250},
			// 			Text:    "Проверка БД АлкоХелп 3",
			// 			OnClicked: func() {
			// 				uri := path.Join(p.app.BaseUrl(), "/v1/checka3")
			// 				p.app.Open(uri)
			// 			},
			// 		},
			// 		dcl.TextLabel{
			// 			// MaxSize: dcl.Size{Width: 400},
			// 			Text: "проверка базы данных алкохелп 3 на предмет дублирования данных",
			// 		},
			// 		dcl.HSpacer{},
			// 	},
			// },

			// основной блок вертикальный пробел для захвата всего пространства внешнего блока
			dcl.VSpacer{},
		},
	}).Create(dcl.NewBuilder(parent)); err != nil {
		return nil, fmt.Errorf("%s %w", modError, err)
	}

	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *MaintainPage) Update() {
}

func (p *MaintainPage) Clear() {
}
