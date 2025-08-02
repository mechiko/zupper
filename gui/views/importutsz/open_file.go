package importutsz

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
		p.checkErrorMessage()
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

	msg := entity.Message{
		Sender: "view:importutsz:openSrcDlg",
		Cmd:    "importttnutsz",
		Model:  nil,
	}
	p.app.Effects().ChanIn() <- msg
}

func (p *ImportTTNPage) proccessTtn(dialog *walk.Dialog, progressBar *walk.ProgressBar) {
	defer dialog.Synchronize(func() {
		dialog.Close(0)
		p.checkErrorMessage()
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
		Sender: "view:importutsz",
		Cmd:    "importttnutsz",
		Model:  nil,
	}
	p.app.Effects().ChanIn() <- msg

}
