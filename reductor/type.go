package reductor

import (
	"sync"

	"go.uber.org/zap"
)

type ModelList map[ModelType]interface{}

type Reductor struct {
	mutex  sync.Mutex
	logger *zap.SugaredLogger
	models ModelList
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
	instance.SetModel(Application, model)
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
