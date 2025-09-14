package application

import (
	"fmt"
	"time"
	"zupper/config"
	"zupper/domain"

	"github.com/mechiko/dbscan"
	"github.com/mechiko/utility"
)

type Application struct {
	model        domain.Model
	Title        string
	Export       string
	Browser      utility.Browser
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
func New(app domain.Apper, repo domain.Repo) (*Application, error) {
	model := &Application{
		model:       domain.Application,
		Title:       "Application Title",
		BrowserList: []string{string(utility.Default), string(utility.Chrome), string(utility.Firefox), string(utility.Yandex), string(utility.Edge)},
	}
	if err := model.ReadState(app, repo); err != nil {
		return nil, fmt.Errorf("model application read state %w", err)
	}
	return model, nil
}

// синхронизирует с приложением в сторону приложения из модели редуктора
func (m *Application) SyncToStore(app domain.Apper) (err error) {
	if err := app.SetOptions("export", m.Export); err != nil {
		return fmt.Errorf("application sync to store: set export failed: %w", err)
	}
	if err := app.SetOptions("browser", m.Browser); err != nil {
		return fmt.Errorf("application sync to store: set browser failed: %w", err)
	}
	return nil
}

// читаем состояние приложения
func (m *Application) ReadState(app domain.Apper, repo domain.Repo) (err error) {
	m.Export = app.Options().Export
	m.Browser = utility.Browser(app.Options().Browser)
	m.Output = app.Options().Output
	m.Host = app.Options().Hostname
	m.Port = app.Options().HostPort
	m.Debug = config.Mode == "development"
	for _, v := range repo.ListDbs() {
		if repo.Is(v) {
			info := repo.Info(v)
			switch v {
			case dbscan.A3:
				m.DbA3Desc = fmt.Sprintf("[%s] %s", info.Driver, info.File)
			case dbscan.Config:
				m.DbConfigDesc = fmt.Sprintf("[%s] %s", info.Driver, info.File)
			case dbscan.TrueZnak:
				m.DbZnakDesc = fmt.Sprintf("[%s] %s", info.Driver, info.File)
			case dbscan.Other:
				m.DbLiteDesc = fmt.Sprintf("[%s] %s", info.Driver, info.File)
			default:
				return fmt.Errorf("application readstate type error %v", v)
			}
		}
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

func (m *Application) Save(app domain.Apper) (err error) {
	if err := app.SaveOptions(); err != nil {
		return fmt.Errorf("application: save options failed: %w", err)
	}
	return nil
}
