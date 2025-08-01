package kontragent

import (
	"zupper/entity"
)

func (p *KontragentListPage) ReductorChangeState(model entity.Model) {
	s := p.app.StartDate()
	e := p.app.EndDate()
	p.start.SetDate(s)
	p.end.SetDate(e)
	p.waitStateLbl.SetText("")
	p.model.SetItems(model.Kontragent.Kontragents)
	p.tv.SetModel(p.model)
}
