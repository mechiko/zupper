package adminka

import (
	"fmt"
	"zupper/domain"
	"zupper/repo"

	"github.com/labstack/echo/v4"
)

const modError = "usecase:adminka"

type adminka struct {
	domain.Apper
	repo *repo.Repository
}

// New -.
func New(a domain.Apper) (*adminka, error) {
	r := repo.GetRepository()
	if r == nil {
		return nil, fmt.Errorf("repo not found")
	}
	adm := &adminka{
		Apper: a,
		repo:  r,
	}
	return adm, nil
}

func (a *adminka) Routes(server *echo.Echo) {
	server.GET("/maintain/statusdb", a.GetReport)
	server.POST("/maintain/statusdbclear", a.PostReport)
}

func (a *adminka) GetReport(c echo.Context) error {
	return a.StatusDb(c)
}

func (a *adminka) PostReport(c echo.Context) error {
	return a.StatusDbClear(c)
}
