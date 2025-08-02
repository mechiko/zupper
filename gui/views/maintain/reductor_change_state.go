package maintain

import (
	"zupper/entity"
)

func (p *MaintainPage) ReductorChangeState(model entity.Model) {
	p.start.SetDate(p.app.StartDate())
	p.end.SetDate(p.app.EndDate())
}
