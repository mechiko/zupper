package views

import (
	"go.uber.org/zap"
)

type IApp interface {
	Logger() *zap.SugaredLogger
	SetBrowser(s string) error
	MessageBox(caption, title string) uintptr
}
