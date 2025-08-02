package kontragent

import (
	"strings"

	"zupper/domain"
	"zupper/gui/types"

	"github.com/mechiko/walk"
)

type KontragentListModel struct {
	walk.TableModelBase
	// walk.SorterBase
	app types.IApp
	// sortColumn int
	// sortOrder  walk.SortOrder
	filter   string
	items    domain.PartnersOriginSlice
	filtered domain.PartnersOriginSlice
}

func NewKontragentListModel(app types.IApp) *KontragentListModel {
	m := new(KontragentListModel)
	m.app = app
	m.filter = ""
	m.ResetRows()
	return m
}

// Called by the TableView from SetModel and every time the model publishes a
// RowsReset event.
func (m *KontragentListModel) RowCount() int {
	return len(m.filtered)
}

func (m *KontragentListModel) Clear() {
	m.items = nil
	m.filtered = nil
}

func (m *KontragentListModel) Items() interface{} {
	return m.filtered
}

// Called by the TableView when it needs the text to display for a given cell.
func (m *KontragentListModel) Value(row, col int) interface{} {
	item := m.filtered[row]
	switch col {
	case 0:
		return string(item.ClientShortName)
	case 1:
		return string(item.ClientInn)
	case 2:
		return string(item.ClientRegID)
	case 3:
		return string(item.ClientCountryCode)
	case 4:
		return string(item.ClientRegionCode)
	case 5:
		return string(item.ClientDescription)
	}
	panic("unexpected col")
}

func (m *KontragentListModel) SetFilter(filter string) {
	if m.filter != filter {
		m.filter = filter
		m.FilteredRows()
	}
}

func (m *KontragentListModel) FilteredRows() {
	m.filtered = nil
	if m.filter != "" {
		for i, item := range m.items {
			if strings.Contains(item.ClientInn, m.filter) {
				m.filtered = append(m.filtered, m.items[i])
				continue
			}
			if strings.Contains(item.ClientRegID, m.filter) {
				m.filtered = append(m.filtered, m.items[i])
				continue
			}
			if strings.Contains(item.ClientDescription, m.filter) {
				m.filtered = append(m.filtered, m.items[i])
				continue
			}
		}
	} else {
		m.filtered = m.items[:]
	}
}

func (m *KontragentListModel) SetItems(src domain.PartnersOriginSlice) {
	m.items = src
	m.ResetRows()
}

func (m *KontragentListModel) ResetRows() {
	// start := m.app.StartDateString()
	// end := m.app.EndDateString()
	// chron, _ := m.app.Repo().DbA3().CatalogPartnerOutgoing(start, end)
	// m.items = chron[:]
	m.FilteredRows()
	// Notify TableView and other interested parties about the reset.
	m.PublishRowsReset()
}
