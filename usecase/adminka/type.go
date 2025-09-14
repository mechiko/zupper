package adminka

import (
	"fmt"
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
	server.GET("/maintain/:report", a.SwitchGetReport)
	server.POST("/maintain/:report", a.SwitchPostReport)
}

func (a *adminka) SwitchGetReport(c echo.Context) error {
	report := c.Param("report")
	switch report {
	case "statusdb":
		a.StatusDb(c)
	case "statusdbclear":
		a.StatusDbClear(c)
	default:
		a.ServerError(c, fmt.Errorf("нет такого отчета"))
	}
	return nil
}

func (a *adminka) SwitchPostReport(c echo.Context) error {
	report := c.Param("report")
	switch report {
	case "statusdbclear":
		a.StatusDbClear(c)
	default:
		a.ServerError(c, fmt.Errorf("нет такого отчета"))
	}
	return nil
}
