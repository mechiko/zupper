package znaktools

import (
	"fmt"
	"net/url"
	"zupper/domain"
	"zupper/domain/models/znakagregate"

	"github.com/mechiko/utility"
	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"
)

func (p *ZnakToolsPage) dclCreate(parent walk.Container, model *znakagregate.ZnakAgregate) error {
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
						Text:     "Выгрузить файл (китай)",
						OnClicked: func() {
							// path.Join cleans slashes and breaks schemes. Use net/url.JoinPath (Go 1.19+) or url.Parse + join on URL.Path
							base := p.BaseUrl()
							uri, jErr := url.JoinPath(base, string(domain.ProdTools))
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
