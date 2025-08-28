package uctemplate

import (
	"bytes"
	"fmt"
	"path/filepath"
	"text/template"
	"zupper/config"

	"github.com/mechiko/utility"
)

func (tt *templateString) tmplMustText(tmpl string, tmplName string, data interface{}, f template.FuncMap) (ss string, err error) {
	defer func() {
		if r := recover(); r != nil {
			ss = ""
			err = fmt.Errorf("panic templateString %v", r)
		}
	}()
	var buf bytes.Buffer

	ftmpl_src := filepath.Join(TemplateSrc, tmplName)
	fncMap := funcMapText
	if f != nil {
		fncMap = f
	}
	if config.Mode == "production" {
		if utility.PathOrFileExists(tmplName) {
			t := template.Must(template.New(tmplName).Funcs(fncMap).ParseFiles(tmplName))
			err = t.ExecuteTemplate(&buf, tmplName, data)
			if err != nil {
				return "", err
			}
		} else {
			t := template.Must(template.New(tmplName).Funcs(fncMap).Parse(tmpl))
			err = t.ExecuteTemplate(&buf, tmplName, data)
			if err != nil {
				return "", err
			}
		}
		return buf.String(), err
	}
	if utility.PathOrFileExists(ftmpl_src) {
		t := template.Must(template.New(tmplName).Funcs(fncMap).ParseFiles(ftmpl_src))
		err = t.ExecuteTemplate(&buf, tmplName, data)
		if err != nil {
			return "", err
		}
	} else if utility.PathOrFileExists(tmplName) {
		t := template.Must(template.New(tmplName).Funcs(fncMap).ParseFiles(tmplName))
		err = t.ExecuteTemplate(&buf, tmplName, data)
		if err != nil {
			return "", err
		}
	} else {
		t := template.Must(template.New(tmplName).Funcs(fncMap).Parse(tmpl))
		err = t.ExecuteTemplate(&buf, tmplName, data)
		if err != nil {
			return "", err
		}
	}
	return buf.String(), err
}
