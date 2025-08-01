package znak

import (
	"fmt"

	"zupper/entity"
)

func (p *ZnakPage) ReductorChangeState(model entity.Model) {
	p.waitStateLbl.SetText("")
	DataPage.GroupOrders = model.Znak.GroupOrders
	DataPage.PackageOrders = model.Znak.PackageOrders
	DataPage.SelectedGroupOrder = model.Znak.SelectedGroupOrder
	DataPage.SelectedPackageOrder = model.Znak.SelectedPackageOrder
	DataPage.ItemPerGroup = model.Znak.ItemPerGroup
	if p.ipsCombo != nil {
		if p.ipsCombo.CurrentIndex() != DataPage.ItemPerGroup {
			p.ipsCombo.SetCurrentIndex(DataPage.ItemPerGroup)
		}
	}

	groupItog := 0
	txt := "выбери заказ для коробок"
	if DataPage.SelectedGroupOrder != nil {
		txt = fmt.Sprintf("Заказ %d GTIN %s %s кол-во %d", DataPage.SelectedGroupOrder.ID, DataPage.SelectedGroupOrder.Gtin, DataPage.SelectedGroupOrder.ProductName, DataPage.SelectedGroupOrder.Quantity)
		groupItog = int(DataPage.SelectedGroupOrder.Quantity)
	}
	p.groupLbl.SetText(txt)
	txt2 := "выбери заказ для бутылок"
	if DataPage.SelectedPackageOrder != nil {
		txt2 = fmt.Sprintf("Заказ %d GTIN %s %s кол-во %d", DataPage.SelectedPackageOrder.ID, DataPage.SelectedPackageOrder.Gtin, DataPage.SelectedPackageOrder.ProductName, DataPage.SelectedPackageOrder.Quantity)
	}
	p.packageLbl.SetText(txt2)
	p.groupItogLbl.SetText(fmt.Sprintf("коробок %d", groupItog))
	if DataPage.SelectedPackageOrder != nil && DataPage.ItemPerGroup > 0 {
		korobok := int(DataPage.SelectedPackageOrder.Quantity) / DataPage.ItemPerGroup
		korobokRemain := int(DataPage.SelectedPackageOrder.Quantity) % DataPage.ItemPerGroup
		p.packageItogLbl.SetText(fmt.Sprintf("хватит на %d коробок с остатком %d;", korobok, korobokRemain))
	}
	if DataPage.SelectedPackageOrder != nil && DataPage.SelectedGroupOrder != nil && DataPage.ItemPerGroup > 0 {
		if DataPage.SelectedPackageOrder.ID != 0 && DataPage.SelectedGroupOrder.ID != 0 && DataPage.ItemPerGroup != 0 {
			p.filePb.SetEnabled(true)
			p.filePb1C.SetEnabled(true)
			p.filePbCsv.SetEnabled(true)
			p.filePbA3.SetEnabled(true)
			p.filePbXml.SetEnabled(true)
		} else {
			p.filePb.SetEnabled(false)
			p.filePb1C.SetEnabled(false)
			p.filePbCsv.SetEnabled(false)
			p.filePbA3.SetEnabled(false)
			p.filePbXml.SetEnabled(false)
		}
	}
	// p.fileLbl.SetText(DataPage.FileName)
}
