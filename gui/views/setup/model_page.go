package setup

import (
	"fmt"
	"zupper/reductor"
)

// создаем новую модель для редуктора, читаем приложение
// и возвращаем структуру
func (p *HomePage) newPageModel() (SetupModel, error) {
	m := &SetupModel{
		Title: "Настройка",
	}
	if err := m.ReadApplication(p.Apper, p.repo); err != nil {
		return SetupModel{}, fmt.Errorf("view:setup ошибка чтения из приложения %w", err)
	}
	return *m, nil
}

// возращаем указатель на модель полученную из редуктора
func (p *HomePage) PageModel() interface{} {
	model := reductor.Instance().Model(p.model)
	return &model
}

// с преобразованием
// если ошибка чтения модели то возвращаем модель из приложения
func (p *HomePage) Model() (*SetupModel, error) {
	if reductor.Instance().IsExistModel(p.model) {
		reductorModel := reductor.Instance().Model(p.model)
		mdl, ok := reductorModel.(SetupModel)
		if ok {
			return &mdl, nil
		} else {
			return nil, fmt.Errorf("view:setup Model другой тип в редукторе %T", mdl)
		}
	}
	return nil, fmt.Errorf("view:setup нет такой модели в редукторе")
}

// сброс модели редуктора для страницы
// и возвращаем указатель
func (p *HomePage) ResetData() interface{} {
	return p.InitData()
}

func (p *HomePage) ResetValidateData() {
}

// инициализируем модель страницы и сохраняет под типом своей модели в редукторе
func (p *HomePage) InitData() interface{} {
	model, err := p.newPageModel()
	if err != nil {
		p.Logger().Errorf("view:home %v", err)
	}
	(&model).ReadApplication(p.Apper, p.repo)
	reductor.Instance().SetModel(p.model, model)
	return &model
}
