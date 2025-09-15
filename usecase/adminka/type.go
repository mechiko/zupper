package adminka

import (
	"zupper/domain"

	"github.com/labstack/echo/v4"
)

const modError = "usecase:adminka"

type adminka struct {
	domain.Apper
}

// New -.
func New(a domain.Apper) *adminka {
	return &adminka{
		Apper: a,
	}
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
