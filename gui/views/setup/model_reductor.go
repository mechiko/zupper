package setup

import (
	"fmt"
	"zupper/config"
	"zupper/domain"
	"zupper/reductor"
	"zupper/repo"
)

type SetupModel struct {
	Model        domain.Model
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
}

// синхронизирует с приложением в сторону приложения из модели редуктора
func (m *SetupModel) Sync(cfg domain.Apper) (err error) {
	modelReductor := reductor.Instance().Model(m.Model)
	model, ok := modelReductor.(SetupModel)
	if !ok {
		return fmt.Errorf("setupModel sync: модель другая %T", modelReductor)
	}
	cfg.SaveOptions("export", model.Export)
	// ...
	return err
}

// берем данные из модели редуктора
func (m *SetupModel) ReadFromModel() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	var model SetupModel
	if reductor.Instance().IsExistModel(m.Model) {
		reductorModel := reductor.Instance().Model(domain.Setup)
		mdl, ok := reductorModel.(SetupModel)
		if !ok {
			return fmt.Errorf("setupModel read: модель другая %T", reductorModel)
		} else {
			model = mdl
		}
	} else {
		return fmt.Errorf("setupModel не инициализирована")
	}
	m.Export = model.Export
	m.Browser = model.Browser
	m.BrowserList = model.BrowserList
	m.Output = model.Output
	m.Debug = model.Debug
	m.Host = model.Host
	m.Port = model.Port
	m.DbLiteDesc = model.DbLiteDesc
	m.DbConfigDesc = model.DbConfigDesc
	m.DbZnakDesc = model.DbZnakDesc
	m.DbA3Desc = model.DbA3Desc
	return nil
}

// читаем состояние приложения
func (m *SetupModel) ReadApplication(app domain.Apper, repo *repo.Repository) (err error) {
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
	return nil
}
