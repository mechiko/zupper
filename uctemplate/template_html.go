package uctemplate

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
	"zupper/config"

	"github.com/mechiko/utility"
)

func (tt *templateString) tmplMustHmtl(tmpl string, tmplName string, data interface{}) (ss string, err error) {
	defer func() {
		if r := recover(); r != nil {
			ss = ""
			err = fmt.Errorf("panic templateString %v", r)
		}
	}()
	var buf bytes.Buffer

	ftmpl_src := filepath.Join(TemplateSrc, tmplName)
	if config.Mode == "production" {
		if utility.PathOrFileExists(tmplName) {
			t := template.Must(template.New(tmplName).Funcs(funcMapHtml).ParseFiles(tmplName))
			err = t.ExecuteTemplate(&buf, tmplName, data)
			if err != nil {
				return "", err
			}
		} else {
			t := template.Must(template.New(tmplName).Funcs(funcMapHtml).Parse(tmpl))
			err = t.ExecuteTemplate(&buf, tmplName, data)
			if err != nil {
				return "", err
			}
		}
		return buf.String(), err
	}
	if utility.PathOrFileExists(ftmpl_src) {
		t := template.Must(template.New(tmplName).Funcs(funcMapHtml).ParseFiles(ftmpl_src))
		err = t.ExecuteTemplate(&buf, tmplName, data)
		if err != nil {
			return "", err
		}
	} else if utility.PathOrFileExists(tmplName) {
		t := template.Must(template.New(tmplName).Funcs(funcMapHtml).ParseFiles(tmplName))
		err = t.ExecuteTemplate(&buf, tmplName, data)
		if err != nil {
			return "", err
		}
	} else {
		t := template.Must(template.New(tmplName).Funcs(funcMapHtml).Parse(tmpl))
		err = t.ExecuteTemplate(&buf, tmplName, data)
		if err != nil {
			return "", err
		}
	}
	return buf.String(), err
}
