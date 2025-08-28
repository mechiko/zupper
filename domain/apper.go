package domain

import (
	"time"
	"zupper/config"

	"go.uber.org/zap"
)

type Apper interface {
	Options() *config.Configuration
	SaveOptions(key string, value interface{}) error
	SaveAllOptions() error
	Logger() *zap.SugaredLogger
	ConfigPath() string
	DefaultDbPath() string
	LogPath() string
	Repo() Repo
	NowDateString() string
	StartDateString() string
	EndDateString() string
	SetStartDate(d time.Time)
	SetEndDate(d time.Time)
	StartDate() time.Time
	EndDate() time.Time
	FsrarID() string
	Pwd() string
}
