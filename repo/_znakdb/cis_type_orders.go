package znakdb

import (
	"fmt"
	"zupper/domain/dbznak"
)

// здесь пинг вызывается в сессии и он не закрывает ее
// и по алгоритму, во всех методах пакета надо закрывать сессию обязательно!
// Единица товара
// Групповая потребительская упаковка
// Набор
func (z *DbZnak) CisTypeOrders(cisType string) (sl []*dbznak.OrderMarkCodes, err error) {
	sess := z.dbSession
	defer sess.Close()

	err = sess.Collection("order_mark_codes").Find("cis_type like ? and archive <> 1", cisType).OrderBy("id desc").All(&sl)
	if err != nil {
		return nil, fmt.Errorf("db:znak upper CisTypeOrders %w", err)
	}
	return sl, nil
}

// if orders, err := znakboil.OrderMarkCodes(qm.Where("cis_type = ? and archive <> 1", cis), qm.OrderBy("id desc")).All(ctx, z.db); err != nil {
// 	if !errors.Is(err, sql.ErrNoRows) {
// 		return ois, fmt.Errorf("%s CisTypeOrders %w", modError, err)
// 	}
// } else {
// 	for i, order := range orders {
// 		newOrder := &domain.OrderInfo{
// 			OrderMarkCode: &domain.OrderMarkCode{},
// 			Guide:         &domain.ProductGuide{},
// 		}
// 		newOrder.OrderMarkCode.ConvertFromZnakSqlite(orders[i])
// 		if guide, err := znakboil.ProductGuides(qm.Where("product_gtin = ?", order.Gtin.String)).One(ctx, z.db); err != nil {
// 			if !errors.Is(err, sql.ErrNoRows) {
// 				z.Logger().Errorf("%s CisTypeOrders %s", modError, err.Error())
// 			}
// 		} else {
// 			newOrder.Guide.ConvertFromZnakSqlite(guide)
// 		}
// 		ois = append(ois, newOrder)
// 	}
// }
