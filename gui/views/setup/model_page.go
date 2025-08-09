package setup

import (
	"fmt"
	"zupper/domain/models/application"
	"zupper/reductor"
	"zupper/utility"
)

// возращаем указатель на модель полученную из редуктора
func (p *SetupPage) PageModel() (interface{}, error) {
	model, err := reductor.Instance().Model(p.model)
	if err != nil {
		return nil, fmt.Errorf("view:setup pagemodel %w", err)
	}
	return model, nil
}

// с преобразованием
// если ошибка чтения модели то возвращаем модель из приложения
func (p *SetupPage) Model() (*application.Application, error) {
	if reductor.Instance().IsExistModel(p.model) {
		reductorModel, err := reductor.Instance().Model(p.model)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		if utility.IsPointer(reductorModel) {
			mdl, ok := reductorModel.(*application.Application)
			if ok {
				return mdl, nil
			} else {
				return nil, fmt.Errorf("view:setup Model другой тип в редукторе %T", mdl)
			}
		}
	}
	return nil, fmt.Errorf("view:setup нет такой модели в редукторе")
}
