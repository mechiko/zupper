package znaktools

import (
	"fmt"
	"zupper/domain/models/znaktool"
	"zupper/reductor"
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
func (p *ZnakToolsPage) Model() (*znaktool.ZnakTools, error) {
	r := reductor.Instance()
	if !r.IsExistModel(p.model) {
		return nil, fmt.Errorf("view:znak нет такой модели в редукторе")
	}
	any, err := r.Model(p.model)
	if err != nil {
		return nil, fmt.Errorf("view:znak get model %w", err)
	}
	mdl, ok := any.(*znaktool.ZnakTools)
	if !ok {
		return nil, fmt.Errorf("view:znak Model другой тип в редукторе %T", any)
	}
	return mdl, nil
}
