package uctemplate

import (
	_ "embed"
	"zupper/domain"
)

//go:embed templates/tmplAdminReport.html
var tmplAdminReport string

type htmlAdminReport struct {
	*domain.AdminReport
	ReportName string
}

func (tt *templateString) PrintAdminReport(adm *domain.AdminReport) (string, error) {
	tmplName := "tmplAdminReport.html"
	report := htmlAdminReport{
		AdminReport: adm,
		ReportName:  "Проверка БД АлкоХелп 3 на дублирование информации",
	}

	// вызов шаблона в него передаем имя шаблона как имя файла шаблона
	if result, err := tt.tmplHmtl(tmplAdminReport, tmplName, report, nil); err != nil {
		return "", err
	} else {
		return result, err
	}
}
