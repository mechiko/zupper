package checkdbg

import (
	"fmt"
	"zupper/repo"

	"go.uber.org/zap"
)

const modError = "pkg:checkdbg"

type Checks struct {
	loger *zap.SugaredLogger
	repo  *repo.Repository
}

func NewChecks(loger *zap.SugaredLogger, repo *repo.Repository) (*Checks, error) {
	// инициализируем REPO
	// TODO изменить получение путей из конфига
	if repo == nil {
		return nil, fmt.Errorf("репозиторий nil")
	}
	return &Checks{
		loger: loger,
		repo:  repo,
	}, nil
}

func (c *Checks) Run() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s Run panic %v", modError, r)
		}
	}()

	if err := c.TestDbA3BuilderGroupMap(); err != nil {
		return err
	}
	if err := c.TestDbA3RawGroupStruct(); err != nil {
		return err
	}
	if err := c.TestDbConfigContact(); err != nil {
		return err
	}
	if err := c.TestDbConfigReleaseMethod(); err != nil {
		return err
	}
	if err := c.TestDbConfigContactWithoutLock(); err != nil {
		return err
	}
	if err := c.TestDbWG(); err != nil {
		return err
	}
	if err := c.TestDbA3CodeApDict(); err != nil {
		return err
	}
	if err := c.TestDbZnakDayUtil(); err != nil {
		return err
	}
	return nil
}
