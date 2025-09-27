package a3

import (
	"fmt"
	"strings"
	"zupper/domain"
)

func (a *DbA3) CodeApMap() (aps map[string]*domain.ApEgais, err error) {
	aps = make(map[string]*domain.ApEgais)
	ap := make([]*domain.ApEgais, 0)
	query := a.dbSession.SQL().
		SelectFrom("ap_egais").Where("id in (select distinct max(ID) from ap_egais group by product_alc_code)")
	err = query.All(&ap)
	if err != nil {
		return nil, fmt.Errorf("repo a3 dict code ap error %w", err)
	}
	for _, cap := range ap {
		code := strings.TrimSpace(cap.ProductAlcCode)
		aps[code] = cap
	}
	return aps, nil
}
