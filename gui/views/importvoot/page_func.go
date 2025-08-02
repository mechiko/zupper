package importvoot

import (
	"path"

	"github.com/mechiko/walk"
)

// func (p *ImportTTNPage) changeIndexBrowser() {
// 	txt := p.browserCB.Text()
// 	// txt := p.browser
// 	p.app.Logger().Debugf("changeIndexBrowser() %s", txt)
// 	if err := p.app.SetBrowser(txt); err != nil {
// 		p.app.Logger().Errorf("changeIndexBrowser %v", err.Error())
// 	}
// }

func (p *ImportTTNPage) changeData() {
	p.app.Logger().Debugf("changeData() %+v", PageData)
}

func (p *ImportTTNPage) openWebApp() {
	uri := path.Join(p.app.BaseUrl(), "/v1/znakreportfull")
	p.app.Open(uri)
}

func (p *ImportTTNPage) UpdateDates() {
}

func (p *ImportTTNPage) openExamImport() {
	uri := path.Join(p.app.BaseUrl(), "/v1/importttn/exam")
	p.app.Open(uri)
}

func (p *ImportTTNPage) openImport() {
	uri := path.Join(p.app.BaseUrl(), "/v1/importttn/import")
	p.app.Open(uri)
}

func (p *ImportTTNPage) openImportCheck() {
	uri := path.Join(p.app.BaseUrl(), "/v1/importttn/check")
	p.app.Open(uri)
}
func bool2int(b bool) walk.CheckState {
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	return walk.CheckState(i)
}

// func (p *ImportTTNPage) SetFifo(b walk.CheckState) {
// 	p.app.Config().Set("import.fifo", (b == 1), true)
// }

// func (p *ImportTTNPage) Fifo() walk.CheckState {
// 	return bool2int(p.app.Configuration().Import.Fifo)
// }

// func (p *ImportTTNPage) SetSplit(b walk.CheckState) {
// 	p.app.Config().Set("import.split", (b == 1), true)
// }

// func (p *ImportTTNPage) Split() (b walk.CheckState) {
// 	return bool2int(p.app.Configuration().Import.Split)
// }

func (p *ImportTTNPage) SetReImport(b walk.CheckState) {
	p.app.Config().Set("import.reimport", (b == 1), true)
}

func (p *ImportTTNPage) ReImport() (b walk.CheckState) {
	return bool2int(p.app.Configuration().Import.ReImport)
}

// func (p *ImportTTNPage) SetIgnoreRest(b walk.CheckState) {
// 	p.app.Config().Set("import.ignorerest", (b == 1), true)
// }

// func (p *ImportTTNPage) IgnoreRest() (b walk.CheckState) {
// 	return bool2int(p.app.Configuration().Import.IgnoreRest)
// }
