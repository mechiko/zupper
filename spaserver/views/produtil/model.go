package produtil

import (
	"fmt"
	"time"
	"zupper/domain"
)

type ProdUtilModel struct {
	Title  string
	Date   time.Time
	model  domain.Model
	errors []error
}

var _ domain.Modeler = (*ProdUtilModel)(nil)

// создаем модель считываем ее состояние и возвращаем указатель
func NewModel(app domain.Apper) (*ProdUtilModel, error) {
	model := &ProdUtilModel{
		model:  domain.ProdTools,
		Title:  "Нанесения сегодня",
		errors: make([]error, 0),
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
	// shallow copy that`s why fields is simple
	dst := *a
	return &dst, nil
}

func (a *ProdUtilModel) Model() domain.Model {
	return a.model
}

func (a *ProdUtilModel) Save(_ domain.Apper) (err error) {
	return nil
}

func (a *ProdUtilModel) Errors() []error {
	out := make([]error, 0, len(a.errors))
	copy(out, a.errors)
	return out
}
