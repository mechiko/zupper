package setup

import (
	"time"
	"zupper/domain"
)

func (p *HomePage) changeIndexBrowser() {
}

func (p *HomePage) Clear() {

}

// обновляет по модели страницу
func (p *HomePage) Update() {
	model, err := p.Model()
	if err != nil {
		p.Logger().Errorf("view:setup update error %v", err)
	}
	if p.lblDbA3 != nil {
		p.lblDbA3.SetText(time.Now().String() + "->" + model.Title)
	}
}

func (p *HomePage) saveConfig() {

}

func (p *HomePage) reloadGui() {
}

func (p *HomePage) SetSendFunc(f func(domain.Model)) {
	p.sendChan = f
}
