package kontragent

import (
	_ "embed"
)

const modError = "gui:view:kontragent"

//go:embed KontragentRegidXml.xml
var KontragentRegidXml string

//go:embed KontragentOborotXml.xml
var KontragentOborotXml string

//go:embed KontragentOborotHtml.html
var KontragentOborotHtml string
