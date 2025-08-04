package reductor

import "zupper/domain"

// вернет модель из мап или nil если запрошенной модели нет
func (rdc *Reductor) Model(page domain.Model) interface{} {
	rdc.mutex.Lock()
	defer rdc.mutex.Unlock()

	if pageModel, ok := rdc.models[page]; ok {
		return pageModel
	}
	return nil
}

// записываем модель по типу енум моделей
func (rdc *Reductor) SetModel(page domain.Model, model interface{}) {
	rdc.mutex.Lock()
	defer rdc.mutex.Unlock()
	if rdc.outStateChan != nil {
		if len(rdc.outStateChan) < cap(rdc.outStateChan) {
			// если канал еще не заполнен то записываем в него тип модели
			rdc.outStateChan <- page
		}
	}
	if rdc.models == nil {
		rdc.models = make(ModelList)
	}
	rdc.models[page] = model
}
