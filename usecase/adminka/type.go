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
	server.GET("/maintain/adminreport", a.AdminReportHtml)
}
