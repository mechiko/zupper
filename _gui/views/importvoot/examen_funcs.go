package importvoot

func (p *ImportTTNPage) clickExamen() {
	PageData.Err = ""
	PageData.Message = ""
	// вызваем прогресс от туда вызывается proccess()
	// p.progressDialog(p.examen)
	// по окончанию вызываем анализ в браузер
	p.openExamImport()
}

// func (p *ImportTTNPage) examen(dialog *walk.Dialog, progressBar *walk.ProgressBar) {
// 	defer dialog.Synchronize(func() {
// 		dialog.Close(0)
// 		p.updateState()
// 	})
// 	f := func(i int) {
// 		dialog.Synchronize(func() {
// 			progressBar.SetValue(i)
// 		})
// 	}
// 	if _, err := p.app.ImportTTN().Examen(f); err != nil {
// 		PageData.Err = err.Error()
// 	}
// }
