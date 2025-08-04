package reductor

import (
	"sync"
	"zupper/domain"

	"go.uber.org/zap"
)

type ModelList map[domain.Model]interface{}

type Reductor struct {
	mutex        sync.Mutex
	logger       *zap.SugaredLogger
	models       ModelList
	outStateChan chan domain.Model
}

var once sync.Once
var instance *Reductor

// создаем singleton и передаем модель по умолчанию reductor.Model("")
func New(logger *zap.SugaredLogger, model interface{}) *Reductor {
	once.Do(func() {
		instance = &Reductor{
			logger: logger,
			models: make(ModelList),
		}
	})
	instance.SetModel(domain.Application, model)
	return instance
}

func Instance() *Reductor {
	if instance == nil {
		panic("reductor instance is nil")
	}
	return instance
}

func (rdc *Reductor) Logger() *zap.SugaredLogger {
	return rdc.logger
}

// если ли в запомненных моделях данная
func (rdc *Reductor) IsExistModel(model domain.Model) bool {
	if _, ok := rdc.models[model]; ok {
		return true
	}
	return false
}

func (rdc *Reductor) SetOutChanState(out chan domain.Model) {
	rdc.outStateChan = out
}
