package app

import (
	"fmt"
	"zupper/domain"
)

type AppModel struct {
	Title   string // заголовок главного окна будет в конфиге не хранится
	License string
	FsrarID string
}

// что храним в конфиге тут прописываем
func (m *AppModel) Sync(cfg domain.Apper) {
	cfg.Options().Application.License = m.License
	cfg.Options().Application.Fsrarid = m.FsrarID
}

// что читаем из конфига тут
func (m *AppModel) Read(cfg domain.Apper) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	m.License = cfg.Options().Application.License
	m.FsrarID = cfg.Options().Application.Fsrarid
	return err
}

func (a *app) InitModel() AppModel {
	model := AppModel{
		Title:   "Application Title",
		License: "",
		FsrarID: "",
	}
	model.Read(a)
	return model
}
