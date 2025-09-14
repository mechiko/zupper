package maintain

import (
	"fmt"
	"zupper/domain/models/application"
	"zupper/reductor"
)

// возращаем указатель на модель полученную из редуктора
func (p *MaintainPage) PageModel() (interface{}, error) {
	model, err := reductor.Instance().Model(p.model)
	if err != nil {
		return nil, fmt.Errorf("view:maintain pagemodel %w", err)
	}
	return model, nil
}

// с преобразованием
// если ошибка чтения модели то возвращаем модель из приложения
func (p *MaintainPage) Model() (*application.Application, error) {
	if reductor.Instance().IsExistModel(p.model) {
		reductorModel, err := reductor.Instance().Model(p.model)
		if err != nil {
			return nil, fmt.Errorf("view:maintain read model: %w", err)
		}
    if mdl, ok := reductorModel.(*application.Application); ok {
      return mdl, nil
    }
    if v, ok := reductorModel.(application.Application); ok {
      return &v, nil
    }
    return nil, fmt.Errorf("view:maintain unexpected model type %T", reductorModel)
	}
	return nil, fmt.Errorf("view:maintain нет такой модели в редукторе")
}
