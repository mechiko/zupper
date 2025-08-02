package views

import (
	"zupper/entity"
	"zupper/utility"
)

// сюда прилетает обновление из редуктора page когда активна страница
func (p *HomePage) ReductorChangeState(model entity.Model) {
	p.disableChange = true
	// ищем браузер в массиве строк и устанавливаем индекс по массиву
	p.Logger().Debugf("%s ChangeState", modError)
	indexBrowser := utility.IndexOf(model.Home.Browser, model.Home.BrowserList)
	p.browserCB.SetCurrentIndex(indexBrowser)
	// p.export.SetText(model.App.Export)
	p.utmhost.SetText(model.Home.UtmHost)
	p.utmport.SetText(model.Home.UtmPort)
	p.lblDbA3.SetText(model.Home.DbA3Desc)
	p.lblDbConfig.SetText(model.Home.DbConfigDesc)
	p.lblDbLite.SetText(model.Home.DbLiteDesc)
	p.lblDbZnak.SetText(model.Home.DbZnakDesc)
	p.disableChange = false
}
