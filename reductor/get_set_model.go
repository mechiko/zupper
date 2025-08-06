package reductor

import (
	"fmt"
	"zupper/domain"
	"zupper/utility"
)

// вернет модель из мап или nil если запрошенной модели нет
// возвращает указатель модели
func (rdc *Reductor) Model(page domain.Model) (interface{}, error) {
	rdc.mutex.RLock()
	defer rdc.mutex.RUnlock()

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
	rdc.models[page] = storeModel.(domain.Modeler)
	// select-based non-blocking send
	if rdc.outStateChan != nil {
		select {
		case rdc.outStateChan <- page:
		default:
			// channel full—drop this update
		}
	}
	return nil
}
