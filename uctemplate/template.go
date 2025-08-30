package uctemplate

type templateString struct {
	// domain.Apper
	layout string
	back   bool
}

func NewTemplate(layout string, back bool) *templateString {
	return &templateString{
		// Apper:  app,
		layout: layout,
		back:   back,
	}
}
