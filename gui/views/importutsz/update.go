package importutsz

func (p *ImportTTNPage) checkErrorMessage() {
	if PageData.Err != "" {
		caption := "Ошибка"
		text := PageData.Err + PageData.Message
		p.app.Logger().Errorf("%s %s", modError, text)
		p.app.MessageBox(caption, text)
	}
	if PageData.Message != "" {
		text := PageData.Message
		p.app.Logger().Infof("%s %s", modError, text)
	}
}
