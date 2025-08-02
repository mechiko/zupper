package reports

import (
	"zupper/entity"
)

func (p *ReportsPage) ReductorChangeState(model entity.Model) {
	p.start.SetDate(p.app.StartDate())
	p.end.SetDate(p.app.EndDate())
}
