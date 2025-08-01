package resource

import (
	"go.uber.org/zap"
)

type IApp interface {
	Logger() *zap.SugaredLogger
}
