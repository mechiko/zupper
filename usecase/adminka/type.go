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
	rp, err := repo.GetRepository()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	adm := &adminka{
		Apper: a,
		repo:  rp,
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
