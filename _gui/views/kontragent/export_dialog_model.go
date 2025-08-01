package kontragent

import (
	"zupper/gui/types"

	"github.com/mechiko/walk"
)

type StateItem struct {
	Title string
	Info  string
	Color string
}

var black walk.Color = walk.RGB(0, 0, 0)
var red walk.Color = walk.RGB(255, 0, 0)

type StateItemSlice []*StateItem

type ListStateModel struct {
	walk.TableModelBase
	// walk.SorterBase
	app types.IApp
	// sortColumn int
	// sortOrder  walk.SortOrder
	filter string
	items  StateItemSlice
}

func NewListStateModel(app types.IApp) *ListStateModel {
	m := new(ListStateModel)
	m.app = app
	m.filter = ""
	m.items = make(StateItemSlice, 0)
	m.ResetRows()
	return m
}

// Called by the TableView from SetModel and every time the model publishes a
// RowsReset event.
func (m *ListStateModel) RowCount() int {
	return len(m.items)
}

func (m *ListStateModel) Clear() {
	m.items = nil
}

func (m *ListStateModel) Items() interface{} {
	return m.items
}

// Called by the TableView when it needs the text to display for a given cell.
func (m *ListStateModel) Value(row, col int) interface{} {
	item := m.items[row]
	switch col {
	case 0:
		return string(item.Title)
	case 1:
		return string(item.Info)
	}
	panic("unexpected col")
}

func (m *ListStateModel) SetItems(src StateItemSlice) {
	m.items = src
	m.ResetRows()
}

func (m *ListStateModel) ResetRows() {
	m.PublishRowsReset()
}

func (m *ListStateModel) AddState(state, colorState string) {
	info := &StateItem{
		Title: state,
		Info:  "",
		Color: colorState,
	}
	m.items = append(m.items, info)
}
