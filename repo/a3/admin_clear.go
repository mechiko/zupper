package a3

import (
	"fmt"
)

var queryRemove = []string{
	`delete from form2_egais where 
  id in (SELECT id FROM form2_egais fe WHERE product_inform_f2_reg_id IN (select product_inform_f2_reg_id from form2_egais GROUP BY product_inform_f2_reg_id  HAVING COUNT(*) > 1)
  AND id NOT IN (select max(id) from form2_egais GROUP BY product_inform_f2_reg_id  HAVING COUNT(*) > 1))
  ;`,
	`delete from form1_egais where 
  id in (SELECT id FROM form1_egais fe WHERE product_inform_f1_reg_id IN (select product_inform_f1_reg_id from form1_egais GROUP BY product_inform_f1_reg_id  HAVING COUNT(*) > 1)
  AND id NOT IN (select max(id) from form1_egais GROUP BY product_inform_f1_reg_id  HAVING COUNT(*) > 1));`,
	`delete from production_form1 where 
  id in (SELECT id FROM production_form1 WHERE id_production_reports IN (select id_production_reports from production_form1 GROUP BY id_production_reports  HAVING COUNT(*) > 1)
  AND id NOT IN (select min(id) from production_form1 GROUP BY id_production_reports  HAVING COUNT(*) > 1));`,
	`delete from ttn_form2 where 
  id in (SELECT id FROM ttn_form2 WHERE id_ttn IN (select id_ttn from ttn_form2 GROUP BY id_ttn  HAVING COUNT(*) > 1)
  AND id NOT IN (select min(id) from ttn_form2 GROUP BY id_ttn  HAVING COUNT(*) > 1));`,
	`update rests_ap_volume
  set product_total_ml_volume = (select top 1 cast(ral.product_quantity as float) * cast(ral.product_capacity as float) * 1000 from rests_ap_local ral where ral.product_inform_f2_reg_id = rests_ap_volume.product_inform_f2_reg_id),
  product_current_ml_volume = (select top 1 cast(ral.product_capacity as float) * 1000 from rests_ap_local ral where ral.product_inform_f2_reg_id = rests_ap_volume.product_inform_f2_reg_id)
  where product_inform_f2_reg_id in (select product_inform_f2_reg_id from rests_ap_volume GROUP BY product_inform_f2_reg_id  HAVING count(*) > 1);`,
	`delete from rests_ap_volume where 
  id in (SELECT id FROM rests_ap_volume fe WHERE product_inform_f2_reg_id IN (select product_inform_f2_reg_id from rests_ap_volume GROUP BY product_inform_f2_reg_id  HAVING COUNT(*) > 1)
  AND id NOT IN (select max(id) from rests_ap_volume GROUP BY product_inform_f2_reg_id  HAVING COUNT(*) > 1));`,
}

func (a *DbA3) AdminReportClear() (err error) {
	for i, sqlStr := range queryRemove {
		result, err := a.dbSession.SQL().Exec(sqlStr)
		if err != nil {
			return fmt.Errorf("%s exec remove script #%d %w", modError, i, err)
		}
		id, _ := result.LastInsertId()
		rows, _ := result.RowsAffected()
		a.logger.Infof("effected id:%d rows:%d", id, rows)
	}
	return nil
}
