package adminka

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a *adminka) ServerError(c echo.Context, err error) error {
	a.Logger().Errorf("%s server error %s", c.Request().RequestURI, err.Error())
	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
}
