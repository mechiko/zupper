package templates

import (
	"embed"
)

//go:embed footer/*.html header/*.html index/*.html prodtools/*.html
var root embed.FS
