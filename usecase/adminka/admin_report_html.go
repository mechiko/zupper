package adminka

import (
	"fmt"
	"net/http"
	"zupper/repo/a3"
	"zupper/uctemplate"

	"github.com/labstack/echo/v4"
	"github.com/mechiko/dbscan"
)

func (a *adminka) StatusDb(c echo.Context) error {
	db, err := a.Repo().Lock(dbscan.A3)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer a.Repo().Unlock(db)

	dbA3, ok := db.(*a3.DbA3)
	if !ok {
		return fmt.Errorf("db type wrong %T", db)
	}

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
	db, err := a.Repo().Lock(dbscan.A3)
	if err != nil {
		return a.ServerError(c, fmt.Errorf("%s %w", modError, err))
	}
	defer a.Repo().Unlock(db)

	dbA3, ok := db.(*a3.DbA3)
	if !ok {
		return a.ServerError(c, fmt.Errorf("%s %w", modError, err))
	}

	err = dbA3.AdminReportClear()
	if err != nil {
		return a.ServerError(c, fmt.Errorf("%s %w", modError, err))
	}
	return c.String(http.StatusOK, "OK")
}
