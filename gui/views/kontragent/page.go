package kontragent

import (
	"path"

	"zupper/entity"
	"zupper/gui/types"

	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"
)

type KontragentListPage struct {
	*walk.Composite
	app          types.IApp
	parent       walk.Form
	tv           *walk.TableView
	model        *KontragentListModel
	boldFont     *walk.Font
	filterLine   *walk.LineEdit
	start        *walk.DateEdit
	end          *walk.DateEdit
	waitStateLbl *walk.Label
}

func NewPage(parent walk.Container, app types.IApp) (pp types.Page, err error) {
	p := new(KontragentListPage)
	p.app = app
	p.parent = parent.Form()
	p.model = NewKontragentListModel(p.app)
	// p.model.ResetRows()
	font := parent.Font()
	p.boldFont, _ = walk.NewFont(font.Family(), font.PointSize(), walk.FontBold)
	if err := (dcl.Composite{
		AssignTo: &p.Composite,
		Name:     "Kontragent",
		Border:   true,
		Layout:   dcl.VBox{SpacingZero: false, Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
		Children: []dcl.Widget{
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, SpacingZero: false,
					Margins: dcl.Margins{Left: 5, Top: 5, Right: 5, Bottom: 5}},
				Children: []dcl.Widget{
					dcl.Label{
						Text: "Периода расчета:",
					},
					dcl.DateEdit{
						Enabled:  true,
						AssignTo: &p.start,
						Format:   "yyyy.MM.dd",
						OnDateChanged: func() {
							p.app.SetStartDate(p.start.Date())
						},
					},
					dcl.DateEdit{
						Enabled:  true,
						AssignTo: &p.end,
						// Format: "yyyy-MM-dd",
						Format: "yyyy.MM.dd",
						OnDateChanged: func() {
							p.app.SetEndDate(p.end.Date())
						},
					},
					dcl.PushButton{
						Text: "Применить",
						OnClicked: func() {
							p.Synchronize(func() {
								p.model.Clear()
								p.tv.SetModel(p.model)
								p.waitStateLbl.SetText("идет запрос контрагентов ...")
							})
							msg := entity.Message{
								Sender: "kontragent.SetPeriod",
								Cmd:    "kontragent",
								Model:  nil,
							}
							p.app.Effects().ChanIn() <- msg
						},
					},
					dcl.Label{
						AssignTo: &p.waitStateLbl,
						Text:     "идет запрос контрагентов ...",
					},
					dcl.HSpacer{},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{MarginsZero: true, Spacing: 5},
				Children: []dcl.Widget{
					dcl.LineEdit{
						AssignTo: &p.filterLine,
						Name:     "filterLine",
						Text:     "",
						MinSize:  dcl.Size{Width: 300},
						OnTextChanged: func() {
							// p.app.Logger().Debugf("text changed %s", p.filterLine.Text())
							// p.model.SetFilter(p.filterLine.Text())
							p.Synchronize(func() {
								p.model.SetFilter(p.filterLine.Text())
								p.app.Logger().Debugf("set model on line change")
								p.tv.SetModel(p.model)
							})
						},
					},
					dcl.HSpacer{},
					dcl.PushButton{
						Text: "Очистить Фильтр",
						OnClicked: func() {
							p.Synchronize(func() { p.filterLine.SetText("") })
						},
						Enabled: dcl.Bind("filterLine.Text != ''"),
					},
					dcl.PushButton{
						Text: "Выгрузить",
						OnClicked: func() {
							p.ExportDialog()
						},
						Enabled: dcl.Bind("tv.HasCurrentItem"),
					},
				},
			},
			dcl.TableView{
				AssignTo:                 &p.tv,
				Name:                     "tv",
				AlternatingRowBG:         true,
				CheckBoxes:               false,
				ColumnsOrderable:         false,
				NotSortableByHeaderClick: true,
				MultiSelection:           false,
				LastColumnStretched:      true,
				Columns: []dcl.TableViewColumn{
					{Title: "Наименование", Alignment: dcl.AlignNear, Width: 300},
					{Title: "ИНН", Alignment: dcl.AlignNear, Width: 130},
					{Title: "ФСРАР ИД", Alignment: dcl.AlignCenter, Width: 130},
					{Title: "Страна", Alignment: dcl.AlignCenter},
					{Title: "Регион", Alignment: dcl.AlignCenter},
					{Title: "Адрес", Alignment: dcl.AlignNear},
				},
				StyleCell: p.styleCell,
				Model:     p.model,
				OnItemActivated: func() {
					k := p.model.filtered[p.tv.CurrentIndex()]
					fsrarId := k.ClientRegID
					uri := path.Join(p.app.BaseUrl(), "/v1/partner/oborot", fsrarId)
					p.app.Open(uri)
				},
			},
		},
	}).Create(dcl.NewBuilder(parent)); err != nil {
		return nil, err
	}

	if err := walk.InitWrapperWindow(p); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *KontragentListPage) styleCell(style *walk.CellStyle) {
}

func (p *KontragentListPage) ExportFunc() {
}

func (p *KontragentListPage) FilterDialog() {
}

func (p *KontragentListPage) UpdateFilter() {
}

func (p *KontragentListPage) Clear() {
	p.model.Clear()
}

func (p *KontragentListPage) ClearFilter() {
}

func (p *KontragentListPage) Update() {
	p.model.ResetRows()
}
