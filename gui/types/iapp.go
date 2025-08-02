package types

import (
	"time"

	"zupper/entity"

	"go.uber.org/zap"
)

type IApp interface {
	Logger() *zap.SugaredLogger
	Config() entity.IConfig
	Configuration() *entity.Configuration

	Pwd() string
	Output() string
	SetOutput(e string)
	SaveConfig() error
	Export() string
	SetExport(s string) error
	Browser() string
	SetBrowser(s string) error
	BaseUrl() string
	OpenDir()
	Open(url string)
	// DumpSql(s string)
	// DumpSqlAppend(s string)
	// DumpSqlClear()
	// SetRepo(entity.Repo)
	Repo() entity.Repo

	InitUtm() error
	// InitDb() error

	// SetGuiService(entity.GuiService)
	// GuiService() entity.GuiService
	Licenser() entity.Licenser
	// SetLicenser(entity.Licenser)
	// SetNeedRestart(bool)
	// NeedRestart() bool
	FsrarID() string
	ScanUTM() bool
	SetScanUTM(bool)

	Reductor() entity.Reductor
	Effects() entity.Effects
	SetReductor(entity.Reductor)
	MessageBox(caption, title string) uintptr
	ImportTTN() entity.ImportTTN

	NowDateString() string
	StartDateString() string
	EndDateString() string
	StartDate() time.Time
	EndDate() time.Time
	SetStartDate(time.Time)
	SetEndDate(time.Time)
}
