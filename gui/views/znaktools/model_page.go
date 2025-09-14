package znaktools

import (
	"fmt"
	"zupper/domain/models/znakagregate"
	"zupper/reductor"

	"github.com/mechiko/utility"
)

// возращаем указатель на модель полученную из редуктора
func (p *ZnakToolsPage) PageModel() (interface{}, error) {
	model, err := reductor.Instance().Model(p.model)
	if err != nil {
		return nil, fmt.Errorf("view:znak pagemodel %w", err)
	}
	return model, nil
}

// с преобразованием
// если ошибка чтения модели то возвращаем модель из приложения
func (p *ZnakToolsPage) Model() (*znakagregate.ZnakAgregate, error) {
	if reductor.Instance().IsExistModel(p.model) {
		reductorModel, err := reductor.Instance().Model(p.model)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		if utility.IsPointer(reductorModel) {
			mdl, ok := reductorModel.(*znakagregate.ZnakAgregate)
			if ok {
				return mdl, nil
			} else {
				return nil, fmt.Errorf("view:znak Model другой тип в редукторе %T", mdl)
			}
		}
	}
	return nil, fmt.Errorf("view:znak нет такой модели в редукторе")
}
