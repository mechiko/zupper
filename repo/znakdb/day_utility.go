package znakdb

import (
	"fmt"
	"zupper/domain"

	"github.com/upper/db/v4"
)

// ищем только товарную группу 18 это пиво
func (z *DbZnak) DayUtilisation(day string) (out []*domain.DayUtilisation, err error) {
	out = make([]*domain.DayUtilisation, 0)
	q := z.dbSession.SQL().
		Select(db.Raw("omc.gtin, omc.id as [order], omu.id as [utilisation],	omu.create_date,	omu.production_date,	omu.quantity,	omu.status,	pg.product_name,	pg.product_alc_code")).
		From("order_mark_utilisation omu").
		Where("omc.template_id = 18 and omu.status in ('Обработан', 'Закрыт') and  SUBSTRING(omu.create_date,1,10) = ?", day).
		Join("order_mark_codes omc").
		On("omc.id = omu.id_order_mark_codes").
		Join("product_guides pg").
		On("pg.product_gtin like omc.gtin")
	err = q.All(&out)
	if err != nil {
		return nil, fmt.Errorf("repo a3 dict code ap error %w", err)
	}
	return out, nil
}
