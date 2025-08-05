package setup

import (
	"fmt"

	"zupper/domain"
	"zupper/domain/models/application"
	"zupper/gui/types"
	"zupper/reductor"
	"zupper/repo"

	"github.com/mechiko/walk"
)

const modError = "gui:view:setup"

type HomePage struct {
	*walk.Composite
	domain.Apper
	disableChange bool
	browserCB     *walk.ComboBox
	model         domain.Model
	repo          *repo.Repository
	sendChan      func(domain.Model)

	utmhost     *walk.LineEdit
	utmport     *walk.LineEdit
	saveconf    *walk.PushButton
	lblDbLite   *walk.Label
	lblDbZnak   *walk.Label
	lblDbConfig *walk.Label
	lblDbA3     *walk.Label
}

// герератор страницы при активации в меню
// берем данные из модели страницы в редукторе и заполняем
func New(parent walk.Container, app domain.Apper, repo *repo.Repository) (pp types.Page, err error) {
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Errorf("%s newHomePage panic %v", modError, r))
		}
	}()
	p := &HomePage{
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
	if err = p.dclCreate(parent, model); err != nil {
		return nil, fmt.Errorf("page setup dcl create %w", err)
	}
	return p, err
}

