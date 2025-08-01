package mainwindow

import (
	"zupper/entity"

	"go.uber.org/zap"
)

type IApp interface {
	Logger() *zap.SugaredLogger
	BaseUrl() string
	Reductor() entity.Reductor
	Effects() entity.Effects
	SetReductor(entity.Reductor)
}
