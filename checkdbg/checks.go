package checkdbg

import (
	"fmt"

	"zupper/domain"
	"zupper/repo"
)

const modError = "pkg:checkdbg"

type Checks struct {
	domain.Apper
	repo *repo.Repository
}

func NewChecks(app domain.Apper, repo *repo.Repository) *Checks {
	return &Checks{
		Apper: app,
		repo:  repo,
	}
}

func (c *Checks) Run() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%s Run panic %v", modError, r)
		}
	}()

	// if err := c.TestUtilityParseCis(); err != nil {
	// 	return err
	// }
	return nil
}
