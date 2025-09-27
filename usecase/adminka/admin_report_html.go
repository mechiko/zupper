package adminka

import (
	"fmt"
	"net/http"
	"zupper/uctemplate"

	"github.com/labstack/echo/v4"
)

func (a *adminka) StatusDb(c echo.Context) error {
	dbA3, err := a.repo.LockA3()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer func() {
		if uerr := a.repo.UnlockA3(dbA3); uerr != nil {
			a.Logger().Errorf("%s unlock error %s", modError, uerr.Error())
		}
	}()

	admReport, err := dbA3.AdminReport()
	if err != nil {
		return a.ServerError(c, fmt.Errorf("%s %w", modError, err))
	}
	strout, err := uctemplate.NewTemplate(a.Options().Layouts.TimeLayoutDay, false).PrintAdminReport(admReport)
	if err != nil {
		return a.ServerError(c, fmt.Errorf("%s %w", modError, err))
	}
	if err := c.HTML(http.StatusOK, strout); err != nil {
		a.Logger().Errorf("%s %s", modError, err.Error())
		return a.ServerError(c, fmt.Errorf("%s %w", modError, err))
	}
	return nil
}

func (a *adminka) StatusDbClear(c echo.Context) error {
	dbA3, err := a.repo.LockA3()
	if err != nil {
		return a.ServerError(c, fmt.Errorf("%s %w", modError, err))
	}
	defer func() {
		if uerr := a.repo.UnlockA3(dbA3); uerr != nil {
			a.Logger().Errorf("%s unlock error %s", modError, uerr.Error())
		}
	}()

	if err := dbA3.AdminReportClear(); err != nil {
		return a.ServerError(c, fmt.Errorf("%s %w", modError, err))
	}
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
