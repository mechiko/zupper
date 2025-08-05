package maintain

import (
	"github.com/mechiko/alcogo4lite/entity"
)

func (p *MaintainPage) ReductorChangeState(model entity.Model) {
	p.start.SetDate(p.app.StartDate())
	p.end.SetDate(p.app.EndDate())
}
