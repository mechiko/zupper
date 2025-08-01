package importutsz

import (
	"fmt"

	"zupper/entity"

	"github.com/mechiko/walk"
)

func (p *ImportTTNPage) ReductorChangeState(model entity.Model) {
	p.lblCountTtn.SetText(model.ImportTTN.CountTTN)
	p.lblFile.SetText(model.ImportTTN.File)
	p.fifoCheckBox.SetCheckState(walk.CheckState(model.ImportTTN.Fifo))
	p.emptyTtnCheckBox.SetCheckState(walk.CheckState(model.ImportTTN.EmptyTtn))
	p.splitCheckBox.SetCheckState(walk.CheckState(model.ImportTTN.Split))
	p.reimportCheckBox.SetCheckState(walk.CheckState(model.ImportTTN.ReImport))
	// если счетчик импортируемых ттн пуст запрещаем кнопки
	if model.ImportTTN.CountTTN == "" {
		p.btnExamen.SetEnabled(false)
		p.btnImportTtn.SetEnabled(false)
		p.btnControlTtn.SetEnabled(false)
	} else {
		if model.ImportTTN.CountProccessError == "" {
			ss := "(ошибок нет)"
			if model.ImportTTN.CountProtocol != "" {
				ss = fmt.Sprintf("(ошибок нет, протокол %s)", model.ImportTTN.CountProtocol)
			}
			p.lblProccessingError.SetText(ss)
			p.btnExamen.SetEnabled(true)
			p.btnImportTtn.SetEnabled(!model.ImportTTN.Imported)
			p.btnControlTtn.SetEnabled(true)
		} else {
			protocolStr := ""
			if model.ImportTTN.CountProtocol != "" {
				protocolStr = fmt.Sprintf(", протокол %s", model.ImportTTN.CountProtocol)
			}
			ss := fmt.Sprintf("(ошибок обработки %s%s)", model.ImportTTN.CountProccessError, protocolStr)
			p.lblProccessingError.SetText(ss)
			p.btnExamen.SetEnabled(true)
			p.btnImportTtn.SetEnabled(false)
			p.btnControlTtn.SetEnabled(false)
		}
	}
}
