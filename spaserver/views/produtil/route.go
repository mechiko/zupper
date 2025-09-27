package produtil

import (
	"net/http"
	"time"
	"zupper/reductor"

	"github.com/labstack/echo/v4"
)

func (t *page) Routes() error {
	t.Echo().GET("/prodtools", t.Index)
	t.Echo().GET("/prodtools/table", t.Table)
	t.Echo().GET("/prodtools/produce/:date", t.Produce)
	t.Echo().GET("/prodtools/:date", t.IndexDate)
	return nil
}

func (t *page) Index(c echo.Context) error {
	data, err := t.PageData()
	if err != nil {
		return t.ServerError(c, err)
	}
	if err := c.Render(http.StatusOK, t.Name(), t.RenderPageModel("index", data)); err != nil {
		return t.ServerError(c, err)
	}
	return nil
}

func (t *page) IndexDate(c echo.Context) error {
	date := c.Param("date")
	if date == "" {
		return t.Index(c)
	}
	data, err := t.PageModel()
	if err != nil {
		return t.ServerError(c, err)
	}
	data.Date, err = time.Parse(t.Options().Layouts.TimeLayoutDay, date)
	if err != nil {
		return t.ServerError(c, err)
	}
	err = reductor.Instance().SetModel(data, false)
	if err != nil {
		return t.ServerError(c, err)
	}
	if err := c.Render(http.StatusOK, t.Name(), t.RenderPageModel("index", data)); err != nil {
		return t.ServerError(c, err)
	}
	return nil
}
