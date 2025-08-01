package znak

import (
	"fmt"

	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"

	"zupper/entity"
	"zupper/gui/types"
	"zupper/usecase/services"
)

const modError = "gui:view:znak"

type ZnakPage struct {
	*walk.Composite
	types.IApp
	parent    walk.Form
	smallFont *walk.Font
	tableFont *walk.Font

	groupLbl       *walk.Label
	packageLbl     *walk.Label
	ipsCombo       *walk.ComboBox
	groupItogLbl   *walk.Label
	packageItogLbl *walk.Label
	fileLbl        *walk.Label
	filePb         *walk.PushButton
	filePbA3       *walk.PushButton
	filePbXml      *walk.PushButton
	filePb1C       *walk.PushButton
	filePbCsv      *walk.PushButton
	waitStateLbl   *walk.Label
}

func NewPage(parent walk.Container, app types.IApp) (pp types.Page, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s NewPage ZnakPage panic %v", modError, r)
		}
	}()
	p := new(ZnakPage)
	p.IApp = app
	p.parent = parent.Form()
	p.smallFont, _ = walk.NewFont("JetBrains Mono", 9, 0)
	p.tableFont, _ = walk.NewFont("JetBrains Mono", 10, walk.FontBold)

	if err := (dcl.Composite{
		Border:   true,
		AssignTo: &p.Composite,
		Name:     "ЧЗ",
		Layout:   dcl.VBox{MarginsZero: false, SpacingZero: false, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
		Children: []dcl.Widget{
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.Label{
						MinSize: dcl.Size{Width: 170},
						Text:    "Штук в группе:",
					},
					dcl.ComboBox{
						AssignTo:  &p.ipsCombo,
						Alignment: dcl.AlignHNearVNear,
						Editable:  false,
						Value:     "",
						Model:     itemPerPage,
						OnCurrentIndexChanged: func() {
							curr := p.ipsCombo.CurrentIndex()
							if curr != DataPage.ItemPerGroup && curr > 0 {
								// if curr > 0 {
								mdl := p.Reductor().Model()
								mdl.Znak.ItemPerGroup = curr
								msg := entity.Message{
									Sender: "homepage.OnCurrentIndexChanged",
									Cmd:    "page",
									Model:  &mdl,
								}
								p.Reductor().ChanIn() <- msg
								// }
							}
						},
					},
					dcl.Label{
						AssignTo: &p.waitStateLbl,
						Text:     "идет расчет аггрегации тары ...",
					},
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.PushButton{
						MinSize:   dcl.Size{Width: 170},
						Text:      "Заказ Коробка",
						OnClicked: p.selectGroupDialog,
					},
					dcl.Label{
						AssignTo: &p.groupLbl,
					},
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.Label{
						AssignTo: &p.groupItogLbl,
						Text:     "",
					},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.PushButton{
						MinSize:   dcl.Size{Width: 170},
						Text:      "Заказ Бутылка",
						OnClicked: p.selectPackageDialog,
					},
					dcl.Label{
						AssignTo: &p.packageLbl,
					},
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.Label{
						AssignTo: &p.packageItogLbl,
						Text:     "",
					},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.PushButton{
						AssignTo: &p.filePb,
						MinSize:  dcl.Size{Width: 170},
						Enabled:  false,
						Text:     "Выгрузить файл (китай)",
						OnClicked: func() {
							if fn, err := services.New(p.IApp).ZnakMergeOrdersSingleColumn(DataPage.SelectedGroupOrder.Guide.ProductName, int(DataPage.SelectedGroupOrder.ID), int(DataPage.SelectedPackageOrder.ID), DataPage.ItemPerGroup); err != nil {
								// if fn, err := services.New(p.IApp).ZnakMergeOrders(DataPage.SelectedGroupOrder.Guide.ProductName, int(DataPage.SelectedGroupOrder.ID), int(DataPage.SelectedPackageOrder.ID), DataPage.ItemPerGroup); err != nil {
								p.IApp.MessageBox("ошибка", err.Error())
								mdl := p.Reductor().Model()
								mdl.Znak.FileName = fn
								msg := entity.Message{
									Sender: "homepage.Выгрузить файл",
									Cmd:    "page",
									Model:  &mdl,
								}
								p.Reductor().ChanIn() <- msg
							} else {
								DataPage.FileName = fn
								p.OpenDir()
								mdl := p.Reductor().Model()
								mdl.Znak.FileName = fn
								msg := entity.Message{
									Sender: "homepage.Выгрузить файл",
									Cmd:    "page",
									Model:  &mdl,
								}
								p.Reductor().ChanIn() <- msg
							}
						},
					},
					dcl.Label{
						AssignTo: &p.fileLbl,
					},
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.PushButton{
						AssignTo: &p.filePbXml,
						MinSize:  dcl.Size{Width: 170},
						Enabled:  false,
						Text:     "Выгрузить файл XML агрегации (ЧЗ)",
						OnClicked: func() {
							if fn, err := services.New(p.IApp).ZnakMergeOrders2XmlAgregation(DataPage.SelectedGroupOrder.Guide.ProductName, int(DataPage.SelectedGroupOrder.ID), int(DataPage.SelectedPackageOrder.ID), DataPage.ItemPerGroup); err != nil {
								p.IApp.MessageBox("ошибка", err.Error())
								mdl := p.Reductor().Model()
								mdl.Znak.FileName = fn
								msg := entity.Message{
									Sender: "homepage.Выгрузить файл",
									Cmd:    "page",
									Model:  &mdl,
								}
								p.Reductor().ChanIn() <- msg
							} else {
								DataPage.FileName = fn
								p.OpenDir()
								mdl := p.Reductor().Model()
								mdl.Znak.FileName = fn
								msg := entity.Message{
									Sender: "homepage.Выгрузить файл",
									Cmd:    "page",
									Model:  &mdl,
								}
								p.Reductor().ChanIn() <- msg
							}
						},
					},
					dcl.Label{
						AssignTo: &p.fileLbl,
					},
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.PushButton{
						AssignTo: &p.filePbA3,
						MinSize:  dcl.Size{Width: 170},
						Enabled:  false,
						Text:     "Выгрузить файл для A3 агрегации (ексель набор)",
						OnClicked: func() {
							if fn, err := services.New(p.IApp).ZnakMergeOrders2ColumnGroupPack(DataPage.SelectedGroupOrder.Guide.ProductName, int(DataPage.SelectedGroupOrder.ID), int(DataPage.SelectedPackageOrder.ID), DataPage.ItemPerGroup); err != nil {
								p.IApp.MessageBox("ошибка", err.Error())
								mdl := p.Reductor().Model()
								mdl.Znak.FileName = fn
								msg := entity.Message{
									Sender: "homepage.Выгрузить файл",
									Cmd:    "page",
									Model:  &mdl,
								}
								p.Reductor().ChanIn() <- msg
							} else {
								DataPage.FileName = fn
								p.OpenDir()
								mdl := p.Reductor().Model()
								mdl.Znak.FileName = fn
								msg := entity.Message{
									Sender: "homepage.Выгрузить файл",
									Cmd:    "page",
									Model:  &mdl,
								}
								p.Reductor().ChanIn() <- msg
							}
						},
					},
					dcl.Label{
						AssignTo: &p.fileLbl,
					},
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.PushButton{
						AssignTo: &p.filePb1C,
						MinSize:  dcl.Size{Width: 170},
						Enabled:  false,
						Text:     "Выгрузить файл для 1с агрегации (ексель)",
						OnClicked: func() {
							if fn, err := services.New(p.IApp).ZnakMergeOrders2Column(DataPage.SelectedGroupOrder.Guide.ProductName, int(DataPage.SelectedGroupOrder.ID), int(DataPage.SelectedPackageOrder.ID), DataPage.ItemPerGroup); err != nil {
								p.IApp.MessageBox("ошибка", err.Error())
								mdl := p.Reductor().Model()
								mdl.Znak.FileName = fn
								msg := entity.Message{
									Sender: "homepage.Выгрузить файл",
									Cmd:    "page",
									Model:  &mdl,
								}
								p.Reductor().ChanIn() <- msg
							} else {
								DataPage.FileName = fn
								p.OpenDir()
								mdl := p.Reductor().Model()
								mdl.Znak.FileName = fn
								msg := entity.Message{
									Sender: "homepage.Выгрузить файл",
									Cmd:    "page",
									Model:  &mdl,
								}
								p.Reductor().ChanIn() <- msg
							}
						},
					},
					dcl.Label{
						AssignTo: &p.fileLbl,
					},
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.PushButton{
						AssignTo: &p.filePbCsv,
						MinSize:  dcl.Size{Width: 170},
						Enabled:  false,
						Text:     "Выгрузить файл для 1с агрегации (CSV)",
						OnClicked: func() {
							if fn, err := services.New(p.IApp).ZnakMergeOrders(DataPage.SelectedGroupOrder.Guide.ProductName, int(DataPage.SelectedGroupOrder.ID), int(DataPage.SelectedPackageOrder.ID), DataPage.ItemPerGroup); err != nil {
								p.IApp.MessageBox("ошибка", err.Error())
								mdl := p.Reductor().Model()
								mdl.Znak.FileName = fn
								msg := entity.Message{
									Sender: "homepage.Выгрузить файл",
									Cmd:    "page",
									Model:  &mdl,
								}
								p.Reductor().ChanIn() <- msg
							} else {
								DataPage.FileName = fn
								p.OpenDir()
								mdl := p.Reductor().Model()
								mdl.Znak.FileName = fn
								msg := entity.Message{
									Sender: "homepage.Выгрузить файл",
									Cmd:    "page",
									Model:  &mdl,
								}
								p.Reductor().ChanIn() <- msg
							}
						},
					},
					dcl.Label{
						AssignTo: &p.fileLbl,
					},
					dcl.HSpacer{},
				},
			},
		},
	}).Create(dcl.NewBuilder(parent)); err != nil {
		return nil, fmt.Errorf("RequestPage %w", err)
	}

	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *ZnakPage) Update() {
	model := p.Reductor().Model()
	msg := entity.Message{
		Sender: "znakpage.Update",
		Cmd:    "page",
		Model:  &model,
	}
	p.Reductor().ChanIn() <- msg
}

func (p *ZnakPage) Clear() {
	p.smallFont.Dispose()
	p.tableFont.Dispose()
}
