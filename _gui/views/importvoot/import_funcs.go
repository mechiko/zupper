package importvoot

import (
	"zupper/entity"

	"github.com/mechiko/walk"
)

func (p *ImportTTNPage) clickImportTtn() {
	PageData.Err = ""
	PageData.Message = ""
	// вызваем прогресс от туда вызывается proccess()
	p.progressDialog(p.importTtn)
	// по окончанию вызываем анализ в браузер
	if PageData.Err == "" {
		p.openImport()
	}
}

func (p *ImportTTNPage) importTtn(dialog *walk.Dialog, progressBar *walk.ProgressBar) {
	defer dialog.Synchronize(func() {
		dialog.Close(0)
		p.updateState()
	})
	f := func(i int) {
		dialog.Synchronize(func() {
			progressBar.SetValue(i)
		})
	}
	if err := p.app.ImportTTN().Proccess(f); err != nil {
		PageData.Err = err.Error()
	}
	msg := entity.Message{
		Sender: "view:importvoot",
		Cmd:    "importttn",
		Model:  nil,
	}
	p.app.Effects().ChanIn() <- msg
}
