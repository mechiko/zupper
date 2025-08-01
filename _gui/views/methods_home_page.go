package views

import (
	"zupper/entity"
)

func (p *HomePage) changeIndexBrowser() {
	if !p.disableChange {
		txt := p.browserCB.Text()
		if err := p.SetBrowser(txt); err != nil {
			p.Logger().Errorf("%s %s", modError, err.Error())
		}
	}
}

func (p *HomePage) Clear() {

}

func (p *HomePage) Update() {

}

func (p *HomePage) saveConfig() {
	host := p.utmhost.Text()
	port := p.utmport.Text()
	if err := p.Config().Set("utmhost", host, true); err != nil {
		p.Logger().Errorf("%s set utmhost in config", modError, err)
	}
	if err := p.Config().Set("utmport", port, true); err != nil {
		p.Logger().Errorf("%s set utmport in config", modError, err)
	}
	// как в первый раз считываем все показатели в состояние
	// _ = utm.New(p.IApp).Ping()
	// _, _ = p.Licenser().HttpReadLicense()
	msg := entity.Message{
		Sender: "homepage.saveConfig",
		Cmd:    "first",
		Model:  nil,
	}
	p.Effects().ChanIn() <- msg
	msg = entity.Message{
		Sender: "homepage.saveConfig",
		Cmd:    "utm",
		Model:  nil,
	}
	p.Effects().ChanIn() <- msg
	msg = entity.Message{
		Sender: "homepage.saveConfig",
		Cmd:    "license",
		Model:  nil,
	}
	p.Effects().ChanIn() <- msg
}

func (p *HomePage) reloadGui() {
	msg := entity.Message{
		Sender: "homepage.reloadGui",
		Cmd:    "reload",
		Model:  nil,
	}
	p.Reductor().ChanIn() <- msg
}
