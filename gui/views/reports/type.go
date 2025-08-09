package reports

import (
	"fmt"

	"zupper/domain"
	"zupper/domain/models/application"
	"zupper/gui/types"
	"zupper/reductor"
	"zupper/repo"

	"github.com/mechiko/walk"
)

const modError = "gui:view:reports"

type ReportsPage struct {
	*walk.Composite
	domain.Apper
	disableChange bool
	browserCB     *walk.ComboBox
	model         domain.Model
	repo          *repo.Repository
	sendChan      func(domain.Model)

	start *walk.DateEdit
	end   *walk.DateEdit
}

// обязательный для реализации интерфейса types.Page
// герератор страницы при активации в меню
// берем данные из модели страницы в редукторе и заполняем
func New(parent walk.Container, app domain.Apper, repo *repo.Repository) (pp types.Page, err error) {
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Errorf("%s newHomePage panic %v", modError, r))
		}
	}()
	p := &ReportsPage{
		Apper:         app,
		repo:          repo,
		disableChange: true,
		model:         domain.Application,
	}
	var model *application.Application
	// инициализируем модель и сохраняем в редукторе
	// если таковой еще нет в нем, предохраняется модель уже созданная и рабочая
	if reductor.Instance().IsExistModel(p.model) {
		model, err = p.Model()
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
	} else {
		return nil, fmt.Errorf("view:setup new нет в редукторе модели %s", p.model)
	}
	// p.start.SetDate(model.StartDate())
	// p.end.SetDate(model.EndDate())
	if err = p.dclCreate(parent, model); err != nil {
		return nil, fmt.Errorf("page setup dcl create %w", err)
	}
	return p, err
}

func (p *ReportsPage) changeIndexBrowser() {
}
