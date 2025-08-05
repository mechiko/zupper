package reductor

import (
	"fmt"
	"zupper/domain"
	"zupper/utility"
)

// вернет модель из мап или nil если запрошенной модели нет
// возвращает указатель модели
func (rdc *Reductor) Model(page domain.Model) (interface{}, error) {
	rdc.mutex.Lock()
	defer rdc.mutex.Unlock()

	if pageModel, ok := rdc.models[page]; ok {
		if !utility.IsPointer(pageModel) {
			return nil, fmt.Errorf("reductor internal error model not pointer")
		}
		return pageModel, nil
	}
	return nil, fmt.Errorf("reductor запрошенной модели нет")
}

// записываем модель по типу енум моделей
// модель должна быть указателем!
// в редукторе модели храним тоже по указателям
func (rdc *Reductor) SetModel(page domain.Model, model domain.Modeler) error {
	rdc.mutex.Lock()
	defer rdc.mutex.Unlock()
	if !utility.IsPointer(model) {
		return fmt.Errorf("reductor: model must be a pointer")
	}
	storeModel, err := model.Copy()
	if err != nil {
		return fmt.Errorf("reductor: само копирования модели %w", err)
	}
	if !utility.IsPointer(storeModel) {
		return fmt.Errorf("reductor: model copy must be a pointer")
	}
	if rdc.models == nil {
		rdc.models = make(ModelList)
	}
	rdc.models[page] = model
	if rdc.outStateChan != nil {
		if len(rdc.outStateChan) < cap(rdc.outStateChan) {
			// если канал еще не заполнен то записываем в него тип модели
			rdc.outStateChan <- page
		}
	}
	return nil
}
