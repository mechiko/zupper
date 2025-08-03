package types

import (
	"zupper/config"

	"go.uber.org/zap"
)

type IApp interface {
	Options() *config.Configuration
	SaveOptions(key string, value interface{}) error
	Logger() *zap.SugaredLogger
	ConfigPath() string
	DbPath() string
	LogPath() string
	BaseUrl() string
}
