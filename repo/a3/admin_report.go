package a3

import (
	"fmt"
	"zupper/domain"
)

func (a *DbA3) AdminReport() (out *domain.AdminReport, err error) {
	out = &domain.AdminReport{}
	sql := `SELECT doc_reg_id as id, COUNT(*) as total FROM production_form1 GROUP BY id_production_reports, doc_reg_id HAVING count(*) > 1;`
	formProd, err := a.partReport(sql)
	if err != nil {
		return nil, fmt.Errorf("production_form1 error %w", err)
	}
	if len(formProd) > 0 {
		out.IsDoubleFormProduce = true
		out.IsDoubleFormProduceRow = formProd
	}
	sql = `select doc_reg_id as id, COUNT(*) as total from import_form1 if2 group by id_import_reports, doc_reg_id HAVING count(*) > 1;`
	formImp, err := a.partReport(sql)
	if err != nil {
		return nil, fmt.Errorf("import_form1 error %w", err)
	}
	if len(formImp) > 0 {
		out.IsDoubleFormImport = true
		out.IsDoubleFormImportRow = formImp
	}
	sql = `select doc_reg_id as id, COUNT(*) as total from ttn_form2 tf GROUP BY id_ttn, doc_reg_id  HAVING count(*) > 1;`
	formTtn, err := a.partReport(sql)
	if err != nil {
		return nil, fmt.Errorf("ttn_form2 error %w", err)
	}
	if len(formTtn) > 0 {
		out.IsDoubleFormTtn = true
		out.IsDoubleFormTtnRow = formTtn
	}

	sql = `select product_inform_f1_reg_id as id, COUNT(*) as total from form1_egais GROUP BY product_inform_f1_reg_id  HAVING count(*) > 1;`
	form1, err := a.partReport(sql)
	if err != nil {
		return nil, fmt.Errorf("form1_egais error %w", err)
	}
	if len(form1) > 0 {
		out.IsDoubleForm1 = true
		out.IsDoubleForm1Row = form1
	}

	sql = `select product_inform_f2_reg_id as id, COUNT(*) as total from form2_egais GROUP BY product_inform_f2_reg_id  HAVING count(*) > 1;`
	form2, err := a.partReport(sql)
	if err != nil {
		return nil, fmt.Errorf("form2_egais error %w", err)
	}
	if len(form2) > 0 {
		out.IsDoubleForm2 = true
		out.IsDoubleForm2Row = form2
	}

	sql = `select DISTINCT tp.product_inform_f1_reg_id as 'id'
	FROM ttn tt join ttn_products tp on tp.id_ttn = tt.id
	where
	tt.doc_date >= '2024.10.01'
	and tp.product_inform_f1_reg_id not in (select product_inform_f1_reg_id from form1_egais)
	;`
	out.AbsentForm1, err = a.absentReport(sql)
	if err != nil {
		return nil, fmt.Errorf("absent product_inform_f1_reg_id error %w", err)
	}

	sql = `select DISTINCT tp.product_inform_f2_reg_id as 'id'
	FROM ttn tt join ttn_products tp on tp.id_ttn = tt.id
	where tt.ttn_type in ('Исходящий', 'Импорт')
	and tt.doc_date >= '2024.10.01'
	and tp.product_inform_f2_reg_id not in (select product_inform_f2_reg_id from form2_egais)
	;`
	out.AbsentForm2, err = a.absentReport(sql)
	if err != nil {
		return nil, fmt.Errorf("absent product_inform_f1_reg_id error %w", err)
	}

	sql = `select DISTINCT tt.consignee_client_reg_id as 'id'
	FROM ttn tt where tt.doc_date >= '2024.10.01' and tt.consignee_client_reg_id not in (select client_reg_id from partners_egais)
	;`
	out.AbsentClient, err = a.absentReport(sql)
	if err != nil {
		return nil, fmt.Errorf("absent product_inform_f1_reg_id error %w", err)
	}

	sql = `select product_inform_f2_reg_id as id, COUNT(*) as total from rests_ap_volume GROUP BY product_inform_f2_reg_id  HAVING count(*) > 1;`
	forms, err := a.partReport(sql)
	if err != nil {
		return nil, fmt.Errorf("rests_ap_volume error %w", err)
	}
	if len(forms) > 0 {
		out.IsDoubleForm2RestVolume = true
		out.IsDoubleForm2RestVolumeRow = forms
	}
	return out, nil
}

func (a *DbA3) partReport(q string) (out domain.FormDoubleSlice, err error) {
	out = make(domain.FormDoubleSlice, 0)
	rows, errRaw := a.dbSession.SQL().Query(q)
	if errRaw != nil {
		return nil, fmt.Errorf("session sql error %w", errRaw)
	}
	iter := a.dbSession.SQL().NewIterator(rows)
	defer iter.Close()
	err = iter.All(&out)
	if err != nil {
		return nil, fmt.Errorf("iterator error %w", err)
	}
	return out, nil
}

func (a *DbA3) absentReport(q string) (out domain.AbsentItemSlice, err error) {
	out = make(domain.AbsentItemSlice, 0)
	rows, errRaw := a.dbSession.SQL().Query(q)
	if errRaw != nil {
		return nil, fmt.Errorf("session sql error %w", errRaw)
	}
	iter := a.dbSession.SQL().NewIterator(rows)
	defer iter.Close()
	err = iter.All(&out)
	if err != nil {
		return nil, fmt.Errorf("iterator error %w", err)
	}
	return out, nil
}
