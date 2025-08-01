package kontragent

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	htm "html/template"
	"text/template"

	"zupper/domain"

	"github.com/friendsofgo/errors"
)

// func (f *kontragentFilter) Submit() {
// 	// если использую DataBinder и на структуру по ссылке то он сам обновляет
// 	core.App().DebugLog().Msg("Submit kontragentFilter")
// }

func (p *KontragentListPage) KontragentExportFunc(kontragent domain.PartnersOrigin) error {
	if err := p.ExportRegidXml(kontragent); err != nil {
		return fmt.Errorf("")
	}
	if err := p.ExportOborotXml(kontragent); err != nil {
		return fmt.Errorf("")
	}
	return nil
}

func (p *KontragentListPage) ExportRegidXml(kontragent domain.PartnersOrigin) error {
	var tpl, tplXml bytes.Buffer
	var err error

	dbSrv := core.App().DbSvc()
	db, err := dbSrv.Db()
	if err != nil {
		fmt.Errorf("")
	}
	defer dbSrv.DbClose(db)

	// файл фирм
	where := ""
	period := core.App().Settings().GetPeriod()
	if !core.App().Settings().GetNoUsePeriod() {
		where = " tt.doc_date >= '" + period.Start() + "' and tt.doc_date <= '" + period.End() + "'"
	}

	prRegid := database.KontragentExportXmlParams{
		Where: where,
		Regid: kontragent.Reg_id,
	}

	t, err := template.New("sql").Parse(database.QueryKontragentProducerShipperPeriod)
	if err != nil {
		return errors.WithMessage(err, "views:[kontragent-export-func.go] 58")
	}
	err = t.Execute(&tpl, prRegid)
	if err != nil {
		return errors.WithMessage(err, "views:[kontragent-export-func.go] 62")
	}
	query := tpl.String()

	core.App().DumpSql(query)
	core.App().LogMemUsage()
	startTime := time.Now()

	sets := core.App().Settings()

	out := []api.KontragentExportRegidItem{}
	if err := db.Select(&out, query); err != nil {
		return errors.WithMessage(err, "database(api):[kontragent-model-excel.go]")
	}

	outTmpl := []api.KontragentExportRegidTemplateItem{}
	for _, row := range out {
		var sh bytes.Buffer

		xml.Escape(&sh, []byte(row.Short_name))
		tmpl := api.KontragentExportRegidTemplateItem{
			Title:     row.Title,
			RegId:     row.Reg_id,
			Country:   row.Country_code,
			ShortName: sh.String(),
			Type:      "",
			Inn:       row.Inn,
			Kpp:       row.Kpp,
			Region:    row.Region_code,
		}
		switch row.Triple {
		case 1:
			tmpl.Type = row.Title
		case 2, 3:
			if strings.Contains(row.Title, "Производитель") {
				tmpl.Type = "Производитель"
			} else if strings.Contains(row.Title, "Импортер") {
				tmpl.Type = "Импортер"
			} else {
				tmpl.Type = row.Title
			}
		default:
			tmpl.Type = row.Title
		}
		outTmpl = append(outTmpl, tmpl)
	}

	core.App().DebugLog().Int("Len(ExportRegid)", len(out)).Send()
	if len(out) == 0 {
		return errors.WithMessage(err, "views:[kontragent-export-func.go] Empty 108")
	}
	tXml, err := template.New("sql").Parse(KontragentRegidXml)
	if err != nil {
		return errors.WithMessage(err, "views:[kontragent-export-func.go]")
	}
	err = tXml.Execute(&tplXml, outTmpl)
	if err != nil {
		return errors.WithMessage(err, "views:[kontragent-export-func.go]")
	}
	// result := tplXml.String()
	fname := p.exportFileName("xml", "Фирмы", kontragent)
	fileName := filepath.Join(sets.GetOutput(), fname)
	if err := p.ToFileXml(fileName, tplXml.String()); err != nil {
		return errors.WithMessage(err, "views:[kontragent-export-func.go]")
	}

	// далее экспорт файла оборота

	core.App().LogMemUsage()
	core.App().DebugLog().Dur("Time Since", time.Since(startTime)).Send()

	return err
}

func (p *KontragentListPage) exportFileName(ext string, t string, kontragent api.KontragentItem) string {
	sets := core.App().Settings()
	short := kontragent.Short_name
	short = strings.ReplaceAll(short, "\"", "")
	short = strings.ReplaceAll(short, "'", "")
	if kontragent.Reg_id != "" {
		short += "_" + kontragent.Reg_id
	}
	if kontragent.Kpp != "" {
		short += "_" + kontragent.Kpp
	}
	short += "_"
	short += sets.GetPeriod().Start() + "_" + sets.GetPeriod().End()
	short += "_" + t + "." + ext
	return short
}

// func (p *KontragentListPage) exportRegidFileName(ext string, short string) string {
// 	sets := core.App().Settings()
// 	short = strings.ReplaceAll(short, "\"", "")
// 	short = strings.ReplaceAll(short, "'", "")
// 	name := ""
// 	if short != "" {
// 		name = short + "_"
// 	}
// 	name += sets.GetPeriod().Start() + "_" + sets.GetPeriod().End()
// 	// name += sets.GetKvartal().String()
// 	name += "_Фирмы." + ext
// 	return name
// }

// func (p *KontragentListPage) exportOborotFileName(ext string, kontragent api.KontragentItem) string {
// 	sets := core.App().Settings()
// 	short := kontragent.Short_name
// 	short = strings.ReplaceAll(short, "\"", "")
// 	short = strings.ReplaceAll(short, "'", "")
// 	name := ""
// 	if short != "" {
// 		name = short + "_"
// 	}
// 	name += sets.GetPeriod().Start() + "_" + sets.GetPeriod().End()
// 	// name += sets.GetKvartal().String()
// 	name += "_Оборот." + ext
// 	return name
// }

// func (p *KontragentListPage) exportHtmlFileName(ext string, kontragent api.KontragentItem) string {
// 	sets := core.App().Settings()
// 	short := kontragent.Short_name
// 	short = strings.ReplaceAll(short, "\"", "")
// 	short = strings.ReplaceAll(short, "'", "")
// 	name := "Отчет_"
// 	if short != "" {
// 		name = short + "_"
// 	}
// 	name += sets.GetPeriod().Start() + "_" + sets.GetPeriod().End()
// 	// name += sets.GetKvartal().String()
// 	name += "_Оборот." + ext
// 	return name
// }

func (p *KontragentListPage) ToFileXml(fname string, s string) error {
	defer utility.RecoverLogWithStack("views:[kontragent-export-func.go] ToRegidXml()")

	file, err := os.OpenFile(fname, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0660)
	if err != nil {
		return errors.WithMessage(err, "views:[kontragent-export-func.go]")
	}
	defer func() error {
		if err := file.Close(); err != nil {
			return err
		}
		return nil
	}()
	_, err = file.WriteString(s)

	return err
}

func (p *KontragentListPage) ExportOborotXml(kontragent api.KontragentItem) error {
	defer utility.RecoverLogWithStack("views:[kontragent-export-func.go] ExportOborotXml()")
	var tpl, tplXml bytes.Buffer
	var err error

	dbSrv := core.App().DbSvc()
	db, err := dbSrv.Db()
	if err != nil {
		return errors.WithMessage(err, "views:[kontragent-export-func.go]")
	}
	defer dbSrv.DbClose(db)

	// файл фирм
	where := ""
	period := core.App().Settings().GetPeriod()
	if !core.App().Settings().GetNoUsePeriod() {
		where = " tt.doc_date >= '" + period.Start() + "' and tt.doc_date <= '" + period.End() + "'"
	}

	prRegid := database.KontragentExportXmlParams{
		Where: where,
		Regid: kontragent.Reg_id,
	}

	t, err := template.New("sql").Parse(database.QueryKontragentOborotPeriod)
	if err != nil {
		return errors.WithMessage(err, "views:[kontragent-export-func.go] 208")
	}
	err = t.Execute(&tpl, prRegid)
	if err != nil {
		return errors.WithMessage(err, "views:[kontragent-export-func.go] 212")
	}
	query := tpl.String()

	core.App().DumpSql(query)
	core.App().LogMemUsage()
	startTime := time.Now()

	sets := core.App().Settings()

	out := []api.KontragentExportOborotItem{}
	if err := db.Select(&out, query); err != nil {
		return errors.WithMessage(err, "database(api):[kontragent-model-excel.go]")
	}
	core.App().DebugLog().Int("Len(Export)", len(out)).Send()
	if len(out) == 0 {
		return errors.WithMessage(err, "views:[kontragent-export-func.go] Empty 235")
	}

	year, month, day := time.Now().Date()
	currentDate := fmt.Sprintf("%02d.%02d.%04d", day, month, year)
	outTmpl := api.KontragentExportOborotTemplate{
		Kontragent:  kontragent,
		CurrentDate: currentDate,
		Start:       sets.GetPeriod().Start(),
		End:         sets.GetPeriod().End(),
		Maps:        make(map[string]api.VcodeOborotTemplate),
	}
	vcodeCount := 1
	for i, v := range out {
		var sh bytes.Buffer
		xml.Escape(&sh, []byte(v.Short))
		if vcode, ok := outTmpl.Maps[v.Vcode]; ok {
			if regid, ok := vcode.Maps[v.RegId]; ok {
				regid.Rows = append(regid.Rows, v)
				vcode.Maps[v.RegId] = regid
			} else {
				reg := api.KontragentExportOborotList{
					Regid: v.RegId,
					Short: sh.String(),
					Inn:   v.Inn,
					Kpp:   v.Kpp,
					Rows:  make([]api.KontragentExportOborotItem, 0),
				}
				reg.Rows = append(reg.Rows, v)
				vcode.Maps[reg.Regid] = reg
				outTmpl.Maps[out[i].Vcode] = vcode
			}
		} else {
			vc := api.VcodeOborotTemplate{
				Vcode: v.Vcode,
				Pn:    vcodeCount,
				Maps:  make(map[string]api.KontragentExportOborotList),
			}
			vcodeCount += 1
			reg := api.KontragentExportOborotList{
				Regid: v.RegId,
				Short: sh.String(),
				Inn:   v.Inn,
				Kpp:   v.Kpp,
				Rows:  make([]api.KontragentExportOborotItem, 0),
			}
			reg.Rows = append(reg.Rows, v)
			vc.Maps[reg.Regid] = reg
			outTmpl.Maps[out[i].Vcode] = vc
		}
	}
	core.App().DebugLog().Str("Current Date", currentDate).Send()

	tXml, err := template.New("sql").Parse(KontragentOborotXml)
	if err != nil {
		return errors.WithMessage(err, "views:[kontragent-export-func.go]")
	}
	err = tXml.Execute(&tplXml, outTmpl)
	if err != nil {
		return errors.WithMessage(err, "views:[kontragent-export-func.go]")
	}

	fname := p.exportFileName("xml", "Обороты", kontragent)
	fileName := filepath.Join(sets.GetOutput(), fname)
	if err := p.ToFileXml(fileName, tplXml.String()); err != nil {
		return errors.WithMessage(err, "views:[kontragent-export-func.go]")
	}

	// далее экспорт файла оборота

	core.App().LogMemUsage()
	core.App().DebugLog().Dur("Time Since", time.Since(startTime)).Send()

	return err
}

func (p *KontragentListPage) ExportOborotHtml(kontragent api.KontragentItem) error {
	defer utility.RecoverLogWithStack("views:[kontragent-export-func.go] ExportOborotXml()")
	var tpl, tplXml bytes.Buffer
	var err error

	dbSrv := core.App().DbSvc()
	db, err := dbSrv.Db()
	if err != nil {
		return errors.WithMessage(err, "views:[kontragent-export-func.go]")
	}
	defer dbSrv.DbClose(db)

	// файл фирм
	where := ""
	period := core.App().Settings().GetPeriod()
	if !core.App().Settings().GetNoUsePeriod() {
		where = " tt.doc_date >= '" + period.Start() + "' and tt.doc_date <= '" + period.End() + "'"
	}

	prRegid := database.KontragentExportXmlParams{
		Where: where,
		Regid: kontragent.Reg_id,
	}

	t, err := template.New("sql").Parse(database.QueryKontragentOborotPeriod)
	if err != nil {
		return errors.WithMessage(err, "views:[kontragent-export-func.go] 208")
	}
	err = t.Execute(&tpl, prRegid)
	if err != nil {
		return errors.WithMessage(err, "views:[kontragent-export-func.go] 212")
	}
	query := tpl.String()

	core.App().DumpSql(query)
	core.App().LogMemUsage()
	startTime := time.Now()

	sets := core.App().Settings()

	out := []api.KontragentExportOborotItem{}
	if err := db.Select(&out, query); err != nil {
		return errors.WithMessage(err, "database(api):[kontragent-model-excel.go]")
	}
	core.App().DebugLog().Int("Len(Export)", len(out)).Send()
	if len(out) == 0 {
		return errors.WithMessage(err, "views:[kontragent-export-func.go] Empty 235")
	}

	year, month, day := time.Now().Date()
	currentDate := fmt.Sprintf("%02d.%02d.%04d", day, month, year)
	outTmpl := api.KontragentExportOborotTemplate{
		BaseUrl:     core.App().Settings().GetBaseUrl(),
		Kontragent:  kontragent,
		CurrentDate: currentDate,
		Start:       sets.GetPeriod().Start(),
		End:         sets.GetPeriod().End(),
		Maps:        make(map[string]api.VcodeOborotTemplate),
	}
	vcodeCount := 1
	var total float64
	total = 0
	for i, v := range out {
		// var sh bytes.Buffer
		// xml.Escape(&sh, []byte(v.Short))
		outVol, err := strconv.ParseFloat(v.Vol, 64)
		total += outVol
		if err != nil {
			return errors.WithMessage(err, "views:[kontragent-export-func.go]")
		}
		if vcode, ok := outTmpl.Maps[v.Vcode]; ok {
			if regid, ok := vcode.Maps[v.RegId]; ok {
				regid.Rows = append(regid.Rows, v)
				vcode.Maps[v.RegId] = regid
			} else {
				reg := api.KontragentExportOborotList{
					Regid: v.RegId,
					Short: v.Short,
					Inn:   v.Inn,
					Kpp:   v.Kpp,
					Rows:  make([]api.KontragentExportOborotItem, 0),
				}
				reg.Rows = append(reg.Rows, v)
				vcode.Maps[reg.Regid] = reg
				outTmpl.Maps[out[i].Vcode] = vcode
			}

		} else {
			vc := api.VcodeOborotTemplate{
				Vcode: v.Vcode,
				Pn:    vcodeCount,
				Maps:  make(map[string]api.KontragentExportOborotList),
			}
			vcodeCount += 1
			reg := api.KontragentExportOborotList{
				Regid: v.RegId,
				Short: v.Short,
				Inn:   v.Inn,
				Kpp:   v.Kpp,
				Rows:  make([]api.KontragentExportOborotItem, 0),
			}
			reg.Rows = append(reg.Rows, v)
			vc.Maps[reg.Regid] = reg
			outTmpl.Maps[out[i].Vcode] = vc
		}
		if vcode, ok := outTmpl.Maps[v.Vcode]; ok {
			vcode.Total += outVol
			outTmpl.Maps[v.Vcode] = vcode
			if regid, ok := vcode.Maps[v.RegId]; ok {
				regid.Total += outVol
				vcode.Maps[v.RegId] = regid
			}
		}

	}
	outTmpl.Total = total
	core.App().DebugLog().Str("Current Date", currentDate).Send()

	funcMap := htm.FuncMap{
		"FormatNumber": func(value float64) string {
			return fmt.Sprintf("%.3f", value)
		},
		"AddOne": func(value int) int {
			return value + 1
		},
	}
	tXml, err := htm.New("sql").Funcs(funcMap).Parse(KontragentOborotHtml)
	if err != nil {
		return errors.WithMessage(err, "views:[kontragent-export-func.go]")
	}
	err = tXml.Execute(&tplXml, outTmpl)
	if err != nil {
		return errors.WithMessage(err, "views:[kontragent-export-func.go]")
	}

	// fname := p.exportHtmlFileName("html", kontragent)
	fname := p.exportFileName("html", "Отчет", kontragent)
	fileName := filepath.Join(sets.GetOutput(), fname)
	if err := p.ToFileXml(fileName, tplXml.String()); err != nil {
		return errors.WithMessage(err, "views:[kontragent-export-func.go]")
	}

	// core.App().Open(fileName, core.App().Config().Browser)
	core.App().LogMemUsage()
	core.App().DebugLog().Dur("Time Since", time.Since(startTime)).Send()

	return err
}
