package domain

import (
	"zupper/config"

	"go.uber.org/zap"
)

type Apper interface {
	Options() *config.Configuration
	SaveOptions(key string, value interface{}) error
	Logger() *zap.SugaredLogger
	ConfigPath() string
	DbPath() string
	LogPath() string
}
