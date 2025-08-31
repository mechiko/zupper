package uctemplate

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"path/filepath"
	"zupper/config"
)

//go:embed templates/index.html
var tmplAdminIndex string

func (tt *templateString) tmplHtml(tmpl string, tmplName string, data interface{}, f template.FuncMap) (ss string, err error) {
	var buf bytes.Buffer

	fncMap := funcMapHtml
	if f != nil {
		fncMap = f
	}
	// общий шаблон
	templ := template.New("tmplindex").Funcs(fncMap)
	// в режиме отладки считываем файл шаблона и обрабатываем его
	if config.Mode == "development" {
		templateFileAbs := filepath.Join(rootPathTemplates(), tmplName)
		templateIndexAbs := filepath.Join(rootPathTemplates(), "index.html")
		strIndexTemplate, err := readFile(templateIndexAbs)
		if err != nil {
			// файл не может быть прочитан выводим ошибку
			return "", fmt.Errorf("uctemplate ошибка файла индекса шаблона %s %w", tmplName, err)
		}
		_, err = templ.New("index").Parse(strIndexTemplate)
		if err != nil {
			return "", fmt.Errorf("uctemplate error parse index.html %s %w", tmplName, err)
		}
		strTemplate, err := readFile(templateFileAbs)
		if err != nil {
			// файл не может быть прочитан выводим ошибку
			return "", fmt.Errorf("uctemplate ошибка файла шаблона %s %w", tmplName, err)
		}
		_, err = templ.New("report").Parse(strTemplate)
		if err != nil {
			return "", fmt.Errorf("uctemplate error parse %s %w", tmplName, err)
		}
		err = templ.ExecuteTemplate(&buf, "index", data)
		if err != nil {
			return "", fmt.Errorf("uctemplate error execute %s %w", tmplName, err)
		}
		return buf.String(), err
	} else {
		_, err = templ.New("index").Parse(tmplAdminIndex)
		if err != nil {
			return "", fmt.Errorf("uctemplate error parse index.html %s %w", tmplName, err)
		}
		_, err = templ.New("report").Parse(tmpl)
		if err != nil {
			return "", fmt.Errorf("uctemplate error parse %s %w", tmplName, err)
		}
		err = templ.ExecuteTemplate(&buf, "index", data)
		if err != nil {
			return "", fmt.Errorf("uctemplate error execute %s %w", tmplName, err)
		}
		return buf.String(), err
	}
}
