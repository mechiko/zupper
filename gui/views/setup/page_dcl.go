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
	defer func() {
		if r := recover(); r != nil {
			p.Logger().Errorf("setup page panic: %v", r)
		}
	}()

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
						Model:    model.BrowserList,
						OnCurrentIndexChanged: func() {
							p.Logger().Debug("browser current index change")
							txt := p.browserCB.Text()
							if txt == "" {
								return
							}
							modelChange, err := p.Model()
							if err != nil {
								p.Logger().Errorf("get model error: %v", err)
								return
							}
							newVal := utility.Browser(txt)
							if modelChange.Browser == newVal {
								return
							}
							modelChange.Browser = newVal
							if err = modelChange.SyncToStore(p); err != nil {
								p.Logger().Errorf("sync browser to store error: %v", err)
								return
							}
							if err = modelChange.Save(p); err != nil {
								p.Logger().Errorf("save model error: %v", err)
								return
							}
							if err = reductor.Instance().SetModel(modelChange, false); err != nil {
								p.Logger().Errorf("set model in reductor error: %v", err)
								return
							}
						},
					},
				},
			},
			dcl.GroupBox{
				Title:  "Конфигурация БД",
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
					dcl.PushButton{
						Text: "Открыть папку выгрузки",
						OnClicked: func() {
							out := p.Options().Output
							if out == "" {
								utility.MessageBox("ошибка", "путь выгрузки не настроен")
								return
							}
							if err := utility.OpenFileInShell(out); err != nil {
								utility.MessageBox("ошибка", err.Error())
							}
						},
					},
					dcl.PushButton{
						Text: "Открыть папку настройки и логов",
						OnClicked: func() {
							cfg := p.ConfigPath()
							if cfg == "" {
								utility.MessageBox("ошибка", "путь к настройкам/логам не настроен")
								return
							}
							if err := utility.OpenFileInShell(cfg); err != nil {
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
