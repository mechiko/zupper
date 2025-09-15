package uctemplate

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"

	"github.com/mechiko/utility"
)

var funcMapHtml = template.FuncMap{
	// The name "inc" is what the function will be called in the template text.
	"inc": func(i int) int {
		return i + 1
	},
	"noescape": func(str string) template.HTML {
		return template.HTML(str)
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
	"dals": func(cap float64, count string) string {
		// capacity, _ := strconv.ParseFloat(cap, 64)
		quantity, _ := strconv.ParseFloat(count, 64)
		return fmt.Sprintf("%0.3f", cap*quantity*0.1)
	},
	"dal": func(cap float64, count int64) string {
		// capacity, _ := strconv.ParseFloat(cap, 64)
		return fmt.Sprintf("%0.3f", cap*float64(count)*0.1)
	},
	"dates": func(d string) template.HTML {
		arr := strings.Split(d, ",")
		out := ""
		if len(arr) > 0 {
			for i := 0; i < len(arr); i++ {
				var dateProduce, dateFinal string
				arrColumn := strings.Split(arr[i], ":")
				utility.Unpack(arrColumn, &dateProduce, &dateFinal)
				// 0123-56-89
				out += fmt.Sprintf("%s.%s.%s - %s.%s.%s<br>", dateProduce[8:], dateProduce[5:7], dateProduce[:4], dateFinal[8:], dateFinal[5:7], dateFinal[:4])
			}
			// return out
			return template.HTML(out)
		} else {
			return template.HTML(d)
		}
	},
	"lists": func(d string) template.HTML {
		d = strings.TrimRight(d, " ")
		d = strings.ReplaceAll(d, " ", "<br>")
		return template.HTML(d)
	},
	"colorCheckLine": func(d1, d2 float64) string {
		d1s := fmt.Sprintf("%.2f", d1)
		d2s := fmt.Sprintf("%.2f", d2)
		if d1s == d2s {
			return "text-black"
		} else {
			return "text-orange-500"
		}
	},
	"apparty": func(s string) string {
		ss := strings.Split(s, ":")
		if len(ss) > 1 {
			return fmt.Sprintf("[%s]", ss[1])
		} else {
			return ""
		}
	},
	"apcode": func(s string) string {
		ss := strings.Split(s, ":")
		if len(ss) > 0 {
			return ss[0]
		}
		return ""
	},
	"apvolume": func(s string) string {
		ss := strings.Split(s, ":")
		if len(ss) > 0 {
			return fmt.Sprintf("%s", ss[1])
		} else {
			return ""
		}
	},
	"balanceFloat64": func(d1, d2 float64) float64 {
		return d1 - d2
	},
	"balanceInt64": func(d1, d2 int64) int64 {
		return d1 - d2
	},
}
