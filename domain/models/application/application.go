package application

import (
	"fmt"
	"time"
	"zupper/config"
	"zupper/domain"
	"zupper/repo"

	"github.com/mechiko/dbscan"
)

type Application struct {
	model        domain.Model
	Title        string
	Export       string
	Browser      string
	BrowserList  []string
	Output       string
	Debug        bool
	Host         string
	Port         string
	DbLiteDesc   string
	DbConfigDesc string
	DbZnakDesc   string
	DbA3Desc     string
	License      string
	FsrarID      string
	startTime    time.Time
	endTime      time.Time
	period       string
}

var _ domain.Modeler = (*Application)(nil)

// создаем модель считываем ее состояние и возвращаем указатель
func New(app domain.Apper, repo *repo.Repository) (*Application, error) {
	model := &Application{
		model: domain.Application,
		Title: "Application Title",
	}
	if err := model.ReadState(app, repo); err != nil {
		return nil, fmt.Errorf("model application read state %w", err)
	}
	return model, nil
}

// синхронизирует с приложением в сторону приложения из модели редуктора
func (m *Application) SyncToStore(app domain.Apper) (err error) {
	app.SaveOptions("export", m.Export)
	// ...
	return err
}

// читаем состояние приложения
func (m *Application) ReadState(app domain.Apper, repo *repo.Repository) (err error) {
	m.Export = app.Options().Export
	m.Browser = app.Options().Browser
	m.Output = app.Options().Output
	m.Host = app.Options().Hostname
	m.Port = app.Options().HostPort
	m.Debug = config.Mode == "development"
	if repo.IsA3() {
		info := repo.Info(dbscan.A3)
		m.DbA3Desc = fmt.Sprintf("[%s] %s", info.Driver, info.File)
	}
	if repo.IsSelf() {
		info := repo.Info(dbscan.Other)
		m.DbLiteDesc = fmt.Sprintf("[%s] %s", info.Driver, info.File)
	}
	if repo.IsA3() {
		info := repo.Info(dbscan.Config)
		m.DbConfigDesc = fmt.Sprintf("[%s] %s", info.Driver, info.File)
	}
	if repo.IsZnak() {
		info := repo.Info(dbscan.TrueZnak)
		m.DbZnakDesc = fmt.Sprintf("[%s] %s", info.Driver, info.File)
	}
	m.License = app.Options().Application.License
	m.FsrarID = app.Options().Application.Fsrarid
	m.InitDateMn()
	return nil
}

func (a *Application) Copy() (interface{}, error) {
	// shallow copy that`s why fields is simple
	dst := *a
	return &dst, nil
}

func (a *Application) Model() domain.Model {
	return a.model
}
