package templates

import (
	"embed"
)

//go:embed header footer home index
var root embed.FS
