package znaktools

import (
	"fmt"
	"net/url"
	"zupper/domain"
	"zupper/domain/models/znaktool"
	"zupper/reductor"

	"github.com/mechiko/utility"
	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"
)

func (p *ZnakToolsPage) dclCreate(parent walk.Container, model *znaktool.ZnakTools) error {
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
					dcl.PushButton{
						AssignTo: &p.filePb,
						MinSize:  dcl.Size{Width: 170},
						Enabled:  true,
						Text:     "Производство по нанесению",
						OnClicked: func() {
							// path.Join cleans slashes and breaks schemes. Use net/url.JoinPath (Go 1.19+) or url.Parse + join on URL.Path
							model, err := p.Model()
							if err != nil {
								p.Logger().Errorf("get model: %v", err.Error())
								return
							}
							base := p.BaseUrl()
							layout := p.Options().Layouts.TimeLayoutDay
							if layout == "" {
								layout = "2006.01.02"
							}
							day := p.date.Format(layout)
							model.Date = p.date
							err = reductor.Instance().SetModel(model, false)
							if err != nil {
								p.Logger().Errorf("get model: %v", err.Error())
								return
							}
							uri, jErr := url.JoinPath(base, string(domain.ProdTools), day)
							if jErr != nil {
								p.Logger().Errorf("build uri: %v", jErr)
								return
							}
							browser := utility.Browser(p.Options().Browser)
							if err := utility.OpenHttpBrowser(uri, browser); err != nil {
								p.Logger().Errorf("open uri error %v", err)
							}
						},
					},
					dcl.DateEdit{
						Enabled:  true,
						AssignTo: &p.dayUtilisation,
						Format:   "yyyy.MM.dd",
						Date:     p.date,
						OnDateChanged: func() {
							p.date = p.dayUtilisation.Date()
							model, err := p.Model()
							if err != nil {
								p.Logger().Errorf("get model: %v", err.Error())
								return
							}
							model.Date = p.date
							err = reductor.Instance().SetModel(model, false)
							if err != nil {
								p.Logger().Errorf("get model: %v", err.Error())
								return
							}
						},
					},
					// dcl.Label{
					// 	AssignTo: &p.fileLbl,
					// 	// TODO: Add specific label field for this button
					// },
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.PushButton{
						AssignTo: &p.pbNumberUtilisation,
						MinSize:  dcl.Size{Width: 170},
						Enabled:  true,
						Text:     "Выгрузить отчет нанесения (Excel)",
						OnClicked: func() {
							p.idUtilisation = int64(p.numberUtilisation.Value())
							if p.idUtilisation <= 0 {
								utility.MessageBox("Ошибка", "номер отчета нанесения должен быть больше нуля")
							}
							if err := p.ExcelUtilisationCodes(p.idUtilisation); err != nil {
								utility.MessageBox("Ошибка", err.Error())
							}
						},
					},
					dcl.NumberEdit{
						Enabled:  true,
						AssignTo: &p.numberUtilisation,
					},
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.PushButton{
						AssignTo: &p.pbNumberOrder,
						MinSize:  dcl.Size{Width: 170},
						Enabled:  true,
						Text:     "Выгрузить заказ (Excel)",
						OnClicked: func() {
							p.idOrder = int64(p.numberOrder.Value())
							if p.idOrder <= 0 {
								utility.MessageBox("Ошибка", "номер заказа должен быть больше нуля")
							}
							if err := p.ExcelOrderCodes(p.idOrder); err != nil {
								utility.MessageBox("Ошибка", err.Error())
							}
						},
					},
					dcl.NumberEdit{
						Enabled:  true,
						AssignTo: &p.numberOrder,
					},
					// dcl.Label{
					// 	AssignTo: &p.fileLbl,
					// 	// TODO: Add specific label field for this button
					// },
					dcl.HSpacer{},
				},
			},
		},
	}).Create(dcl.NewBuilder(parent)); err != nil {
		return fmt.Errorf("gui:view znak %w", err)
	}

	return nil
}
