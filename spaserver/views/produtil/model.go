package produtil

import (
	"fmt"
	"time"
	"zupper/domain"
)

type ProdUtilModel struct {
	Title    string
	Date     time.Time
	model    domain.Model
	Table    []*domain.DayUtilisation
	MapTable map[string]map[string]int
	Reports  []*PrdReport
	errors   []error
}

var _ domain.Modeler = (*ProdUtilModel)(nil)

// создаем модель считываем ее состояние и возвращаем указатель
func NewModel(app domain.Apper) (*ProdUtilModel, error) {
	model := &ProdUtilModel{
		model:  domain.ProdTools,
		Title:  "Нанесения сегодня",
		errors: make([]error, 0),
		Date:   time.Now(),
	}
	if err := model.ReadState(app); err != nil {
		return nil, fmt.Errorf("model prodtools read state %w", err)
	}
	return model, nil
}

// синхронизирует с приложением в сторону приложения из модели редуктора
func (m *ProdUtilModel) SyncToStore(app domain.Apper) (err error) {
	return err
}

// читаем состояние приложения
func (m *ProdUtilModel) ReadState(app domain.Apper) (err error) {
	return nil
}

func (a *ProdUtilModel) Copy() (interface{}, error) {
	dst := *a
	// deep-copy maps
	if a.MapTable != nil {
		dst.MapTable = make(map[string]map[string]int, len(a.MapTable))
		for day, m := range a.MapTable {
			inner := make(map[string]int, len(m))
			for k, v := range m {
				inner[k] = v
			}
			dst.MapTable[day] = inner
		}
	}
	// copy slice headers (new backing arrays)
	if a.Table != nil {
		dst.Table = append([]*domain.DayUtilisation(nil), a.Table...)
	}
	if a.Reports != nil {
		dst.Reports = append([]*PrdReport(nil), a.Reports...)
	}
	if a.errors != nil {
		dst.errors = append([]error(nil), a.errors...)
	}
	return &dst, nil
}

func (a *ProdUtilModel) Model() domain.Model {
	return a.model
}

func (a *ProdUtilModel) Save(_ domain.Apper) (err error) {
	return nil
}

func (a *ProdUtilModel) Errors() []error {
	out := make([]error, len(a.errors))
	copy(out, a.errors)
	return out
}
