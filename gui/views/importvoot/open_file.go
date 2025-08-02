package importvoot

import (
	"path/filepath"

	"zupper/entity"
	"zupper/utility"

	"github.com/mechiko/walk"
)

func (p *ImportTTNPage) openSrcDlg() {
	defer func() {
		if r := recover(); r != nil {
			p.app.Logger().Errorf("%s panic openSrcDlg %v", modError, r)
		}
	}()
	defer func() {
		p.updateState()
	}()
	PageData.Err = ""
	PageData.Message = ""

	dlg := walk.FileDialog{
		InitialDirPath: PageData.Input,
	}
	dlg.Filter = "*.xml"
	dlg.Title = "Выберите файл .xml"
	ok, err := dlg.ShowOpen(p.form)
	if err != nil {
		PageData.Err = err.Error()
		return
	}
	if !ok {
		PageData.Err = "ошибка открытия диалога"
		return
	}
	if utility.PathOrFileExists(dlg.FilePath) {
		PageData.File = dlg.FilePath
		PageData.FileName = filepath.Base(PageData.File)
	} else {
		PageData.Err = "ошибка файл недоступен"
		return
	}
	// if err := p.app.ImportTTN().Init("", PageData.File); err != nil {
	// 	PageData.Err = err.Error()
	// 	return
	// }
	p.progressDialog(p.proccessTtn)
	PageData.Message = "файл открыт"
	// PageData.CountTtn = p.app.ImportTTN().CountString()

	msg := entity.Message{
		Sender: "view:importvoot:openSrcDlg",
		Cmd:    "importttn",
		Model:  nil,
	}
	p.app.Effects().ChanIn() <- msg
}

func (p *ImportTTNPage) proccessTtn(dialog *walk.Dialog, progressBar *walk.ProgressBar) {
	defer dialog.Synchronize(func() {
		dialog.Close(0)
		p.updateState()
	})
	f := func(i int) {
		dialog.Synchronize(func() {
			progressBar.SetValue(i)
		})
	}
	if err := p.app.ImportTTN().Init("", PageData.File, f); err != nil {
		PageData.Err = err.Error()
	}
	msg := entity.Message{
		Sender: "view:importvoot",
		Cmd:    "importttn",
		Model:  nil,
	}
	p.app.Effects().ChanIn() <- msg

}
