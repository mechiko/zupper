package setup

import (
	"fmt"

	"zupper/domain"
	"zupper/gui/types"

	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"
)

const modError = "gui:view:home"

type HomePage struct {
	*walk.Composite
	domain.Apper
	disableChange bool
	browserCB     *walk.ComboBox

	utmhost     *walk.LineEdit
	utmport     *walk.LineEdit
	saveconf    *walk.PushButton
	lblDbLite   *walk.Label
	lblDbZnak   *walk.Label
	lblDbConfig *walk.Label
	lblDbA3     *walk.Label
}

// герератор страницы при активации в меню
func New(parent walk.Container, app domain.Apper) (pp types.Page, err error) {
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Errorf("%s newHomePage panic %v", modError, r))
		}
	}()
	p := &HomePage{
		Apper:         app,
		disableChange: true,
	}
	if err := (dcl.Composite{
		AssignTo: &p.Composite,
		// DataBinder: dcl.DataBinder{
		// 	AssignTo:            &p.db,
		// 	Name:                "Hd",
		// 	ErrorPresenter:      dcl.ToolTipErrorPresenter{},
		// 	OnDataSourceChanged: p.changeData,
		// },
		Name: "homePage",
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
						// Background: dcl.SolidColorBrush{},
						// Background:            dcl.SolidColorBrush{Color: walk.RGB(0x34, 0x82, 0xeb)},
						Editable:              false,
						Value:                 "",
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
						AssignTo:  &p.Composite,
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
								Text:     "",
							},
							dcl.HSpacer{},
						},
					},
					dcl.Composite{
						AssignTo:  &p.Composite,
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
								Text:     "",
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
						AssignTo:  &p.Composite,
						Layout:    dcl.HBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 0, Top: 0, Right: 0, Bottom: 0}},
						Border:    false,
						Alignment: dcl.Alignment2D(walk.AlignHNearVNear),
						Children: []dcl.Widget{
							dcl.Label{
								Text: "БД Config:",
							},
							dcl.Label{
								AssignTo: &p.lblDbConfig,
								Text:     "",
							},
							dcl.HSpacer{},
						},
					},
					dcl.Composite{
						AssignTo:  &p.Composite,
						Layout:    dcl.HBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 0, Top: 0, Right: 0, Bottom: 0}},
						Border:    false,
						Alignment: dcl.Alignment2D(walk.AlignHNearVNear),
						Children: []dcl.Widget{
							dcl.Label{
								Text: "БД A3:",
							},
							dcl.Label{
								AssignTo: &p.lblDbA3,
								Text:     "",
							},
							dcl.HSpacer{},
						},
					},
					dcl.Composite{
						AssignTo:  &p.Composite,
						Layout:    dcl.HBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 0, Top: 0, Right: 0, Bottom: 0}},
						Border:    false,
						Alignment: dcl.Alignment2D(walk.AlignHNearVNear),
						Children: []dcl.Widget{
							dcl.Label{
								Text: "БД ЧЗ:",
							},
							dcl.Label{
								AssignTo: &p.lblDbZnak,
								Text:     "",
							},
							dcl.HSpacer{},
						},
					},
					dcl.Composite{
						AssignTo:  &p.Composite,
						Layout:    dcl.HBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 0, Top: 0, Right: 0, Bottom: 0}},
						Border:    false,
						Alignment: dcl.Alignment2D(walk.AlignHNearVNear),
						Children: []dcl.Widget{
							dcl.Label{
								Text: "БД AlcoGo4Lite:",
							},
							dcl.Label{
								AssignTo: &p.lblDbLite,
								Text:     "",
							},
							dcl.HSpacer{},
						},
					},
				}},

			dcl.Composite{
				Border: false,
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: true, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 0}},
				Children: []dcl.Widget{
					dcl.Composite{
						Layout:  dcl.VBox{MarginsZero: true, SpacingZero: false, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
						MinSize: dcl.Size{Width: 300},
						Children: []dcl.Widget{
							dcl.PushButton{
								AssignTo: &p.saveconf,
								// ColumnSpan: 2,
								Text: "Обновить конфигурацию",
								// OnClicked: p.saveConfig,
								OnClicked: func() {
									// msg := entity.Message{
									// 	Sender: "homepage.reloadGui",
									// 	Cmd:    "reload",
									// 	Model:  nil,
									// }
									// p.Reductor().ChanIn() <- msg
								},
							},
							dcl.PushButton{
								Text: "Открыть папку выгрузки",
								OnClicked: func() {
									// p.OpenDir()
								},
							},
						}}}},
			dcl.VSpacer{},
		},
	}).Create(dcl.NewBuilder(parent)); err != nil {
		return nil, fmt.Errorf("%s %w", modError, err)
	}
	// if err := walk.InitWrapperWindow(p); err != nil {
	// 	return nil, fmt.Errorf("%s %w", modError, err)
	// }
	return p, err
}
