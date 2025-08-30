package uctemplate

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"
	"zupper/config"
)

func (tt *templateString) tmplText(tmpl string, tmplName string, data interface{}, f template.FuncMap) (ss string, err error) {
	var buf bytes.Buffer

	templateFileAbs := filepath.Join(rootPathTemplates(), tmplName)
	fncMap := funcMapText
	if f != nil {
		fncMap = f
	}
	// в режиме отладки считываем файл шаблона и обрабатываем его
	if config.Mode == "development" {
		strTemplate, err := readFile(templateFileAbs)
		if err != nil {
			// файл не может быть прочитан выводим ошибку
			return "", fmt.Errorf("uctemplate ошибка файла шаблона %s %w", tmplName, err)
		}
		t, err := template.New(tmplName).Funcs(fncMap).Parse(strTemplate)
		if err != nil {
			return "", fmt.Errorf("uctemplate error parse %s %w", tmplName, err)
		}
		err = t.ExecuteTemplate(&buf, tmplName, data)
		if err != nil {
			return "", fmt.Errorf("uctemplate error execute %s %w", tmplName, err)
		}
		return buf.String(), err
	} else {
		t, err := template.New(tmplName).Funcs(fncMap).Parse(tmpl)
		if err != nil {
			return "", fmt.Errorf("uctemplate error parse %s %w", tmplName, err)
		}
		err = t.ExecuteTemplate(&buf, tmplName, data)
		if err != nil {
			return "", fmt.Errorf("uctemplate error execute %s %w", tmplName, err)
		}
		return buf.String(), err
	}
}
