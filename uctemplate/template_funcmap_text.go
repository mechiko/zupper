package uctemplate

import (
	"bytes"
	"encoding/xml"
	"strings"
	"text/template"
)

var funcMapText = template.FuncMap{
	// The name "inc" is what the function will be called in the template text.
	"inc": func(i int) int {
		return i + 1
	},
	"volume": func(s string) string {
		ss := strings.Split(s, ":")
		if len(ss) > 1 {
			return ss[1]
		}
		return ""
	},
	"vcode": func(s string) string {
		ss := strings.Split(s, ":")
		if len(ss) > 0 {
			return ss[0]
		}
		return ""
	},
	"escape": func(s string) string {
		var sh bytes.Buffer
		xml.Escape(&sh, []byte(s))
		return sh.String()
	},
}
