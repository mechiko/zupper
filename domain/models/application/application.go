package application

import (
	"fmt"
	"zupper/config"
	"zupper/domain"
	"zupper/repo"
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
}

var _ domain.Modeler = (*Application)(nil)

// создаем модель считываем ее состояние и возвращаем указатель
func NewModelApplication(app domain.Apper, repo *repo.Repository) (*Application, error) {
	model := Application{
		model: domain.Application,
		Title: "Application Title",
	}
	if err := model.ReadState(app, repo); err != nil {
		return nil, fmt.Errorf("model application read state %w", err)
	}
	return &model, nil
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
	m.DbLiteDesc = repo.Self().DbPath()
	m.DbConfigDesc = repo.ConfigPath()
	m.DbZnakDesc = repo.ZnakDB().DbPath()
	m.DbA3Desc = repo.A3DB().DbPath()
	m.License = app.Options().Application.License
	m.FsrarID = app.Options().Application.Fsrarid
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
