package znakpacks

import (
	"github.com/mechiko/alcogo4lite/entity"
	"github.com/mechiko/alcogo4lite/gui/views/znak/models"
	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"
)

func (p *ZnakPage) selectPackageDialog() {
	var dlg *walk.Dialog
	var tv *walk.TableView
	var search *walk.LineEdit
	// var model *models.OrderModel

	model := models.NewModel(p.IApp)
	model.SetItems(DataPage.PackageOrders)
	if err := (dcl.Dialog{
		AssignTo:   &dlg,
		Visible:    true,
		Title:      "Выбор заказа КМ бутылки",
		Borderless: false,
		MinSize:    dcl.Size{Width: 1000, Height: 500},
		Layout:     dcl.VBox{MarginsZero: true},
		Children: []dcl.Widget{
			dcl.LineEdit{
				AssignTo: &search,
				Text:     "",
				OnTextChanged: func() {
					model.Filter(search.Text())
				},
			},
			dcl.TableView{
				AssignTo: &tv,
				// MaxSize:          dcl.Size{Width: 300, Height: 300},
				// MinSize:          dcl.Size{Width: 300, Height: 300},
				AlternatingRowBG:    true,
				CheckBoxes:          false,
				ColumnsOrderable:    false,
				LastColumnStretched: true,
				MultiSelection:      false,
				Columns: []dcl.TableViewColumn{
					{Title: "Заказ", Width: 90},
					{Title: "GTIN", Width: 150},
					{Title: "Наименование", Width: 600},
					{Title: "Кол-во"},
				},
				Model: model,
				OnItemActivated: func() {
					curr := tv.CurrentIndex()
					if curr <= model.RowCount() {
						mdl := p.Reductor().Model()
						mdl.Znak.SelectedPackageOrder = model.Item(curr)
						msg := entity.Message{
							Sender: "homepage.selectPackageDialog",
							Cmd:    "page",
							Model:  &mdl,
						}
						p.Reductor().ChanIn() <- msg
					}
					dlg.Accept()
				},
				// OnSelectedIndexesChanged: func() {
				// 	fmt.Printf("SelectedIndexes: %v\n", tv.SelectedIndexes())
				// 	dlg.Accept()
				// },
			},
		},
	}).Create(p.Form()); err != nil {
		// Hd.Err = err.Error()
		p.Logger().Errorf("%s %s", modError, err.Error())
	}

	dlg.SetFont(p.Font())
	dlg.SetIcon(p.Form().Icon())
	tv.SetModel(model)
	model.ResetRows()
	// search.SetText(Hd.ActModel.GetFilter())

	dlg.Run()
	p.Update()
	dlg.Dispose()
}
