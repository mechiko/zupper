package models

import (
	"strings"

	"zupper/domain"
	"zupper/gui/types"

	"github.com/mechiko/walk"
)

type OrderModel struct {
	walk.TableModelBase
	app       types.IApp
	filter    string
	items     []*domain.OrderInfo
	itemsShow []*domain.OrderInfo
}

func NewModel(a types.IApp) *OrderModel {
	m := &OrderModel{
		app: a,
	}
	return m
}

// Called by the TableView from SetModel and every time the model publishes a
// RowsReset event.
func (m *OrderModel) RowCount() int {
	return len(m.itemsShow)
}

func (m *OrderModel) SetItems(src domain.OrderInfoSlice) {
	m.items = src
	m.ResetRows()
}

// Called by the TableView when it needs the text to display for a given cell.
func (m *OrderModel) Value(row, col int) interface{} {
	if row >= len(m.itemsShow) {
		return ""
	}
	item := m.itemsShow[row]

	switch col {
	case 0:
		// return item.id
		return item.ID
	case 1:
		// return item.number
		return item.Gtin
	case 2:
		// return item.date
		return item.ProductName
	case 3:
		// return item.kontragent
		return item.Quantity
	}

	return "unexpected col"
}

func (m *OrderModel) ResetRows() {
	// m.itemsShow = make(domain.OrderInfoSlice, 0)
	// for i := range m.items {
	// 	m.itemsShow = append(m.itemsShow, newRecord(m.items[i]))
	// }
	m.Filter(m.filter)
	m.PublishRowsReset()
}

// enbed структуру в новую
// func newRecord(pr *domain.OrderInfo) *domain.OrderInfo {
// 	do := *pr
// 	return &do
// }

func (m *OrderModel) GetFilter() string {
	return m.filter
}

func (m *OrderModel) Item(i int) *domain.OrderInfo {
	return m.itemsShow[i]
}

func (m *OrderModel) Filter(search string) {
	// if search == "" {
	// 	return
	// }
	m.itemsShow = make(domain.OrderInfoSlice, 0)
	for i := range m.items {
		if containsStr(search, m.items[i]) {
			m.itemsShow = append(m.itemsShow, m.items[i])
		}
	}
	m.PublishRowsReset()
	m.PublishRowsChanged(0, len(m.itemsShow))
}

func containsStr(str string, r *domain.OrderInfo) bool {
	if strings.Contains(r.Gtin, str) {
		return true
	}
	if strings.Contains(r.ProductName, str) {
		return true
	}
	return false
}
