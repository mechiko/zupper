package uctemplate

import (
	"zupper/domain"
)

type templateString struct {
	domain.Apper
	layout string
	back   bool
}

// путь до шаблонов если нужен режим отладки и редактирования на ходу
const TemplateSrc = `E:\src\goproj\!!alcogo4lite\src\usecase\services\uctemplate`

// const TemplateSrc = `C:\!src\alcogo4lite\src\usecase\services\uctemplate`

func NewTemplate(app domain.Apper, layout string, back bool) *templateString {
	if layout == "" {
		layout = app.Options().Layouts.TimeLayoutDay
	}
	return &templateString{
		Apper:  app,
		layout: layout,
		back:   back,
	}
}
