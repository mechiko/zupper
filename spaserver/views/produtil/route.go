package produtil

import (
	"net/http"
	"zupper/domain"

	"github.com/labstack/echo/v4"
)

func (t *page) Routes() error {
	prodUtil := t.Echo().Group("/" + string(domain.ProdTools))
	prodUtil.GET("", t.Index)
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
