package importvoot

func (p *ImportTTNPage) updateState() {
	// p.db.Reset()
	p.Synchronize(func() {
		defer func() {
			if r := recover(); r != nil {
				p.app.Logger().Errorf("%s panic %v", modError, r)
			}
		}()
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
		// if p.app.ImportTTN().Fifo() {
		// 	p.fifoCheckBox.SetCheckState(1)
		// } else {
		// 	p.fifoCheckBox.SetCheckState(0)
		// }
		// if p.app.ImportTTN().Split() {
		// 	p.splitCheckBox.SetCheckState(1)
		// } else {
		// 	p.splitCheckBox.SetCheckState(0)
		// }
		if p.app.ImportTTN().ReImport() {
			p.reimportCheckBox.SetCheckState(1)
		} else {
			p.reimportCheckBox.SetCheckState(0)
		}
		if p.app.ImportTTN().EmptyTtn() {
			p.emptyTtnCheckBox.SetCheckState(1)
		} else {
			p.emptyTtnCheckBox.SetCheckState(0)
		}
		if PageData.Err != "" {
			// p.btnExamen.SetEnabled(false)
			// p.btnImportTtn.SetEnabled(false)
			return
		}
		if PageData.CountTtn != "" {
			p.lblCountTtn.SetText(PageData.CountTtn)
			p.lblFile.SetText(PageData.FileName)
			// p.btnImportTtn.SetEnabled(true)
			// p.btnExamen.SetEnabled(true)
			p.lblCountTtn.SetText(PageData.CountTtn)
			p.lblFile.SetText(PageData.FileName)
		}
		//  else {
		// 	p.btnImportTtn.SetEnabled(false)
		// 	p.btnExamen.SetEnabled(false)
		// }
		// p.filterBottlingFld.SetText(p.app.ImportTTN().FilterBolling())
		// p.startBottlingFld.SetText(p.app.ImportTTN().StartBotlling())
	})
}
