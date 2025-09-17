package produtil

import (
	"fmt"
	"zupper/repo"

	"github.com/labstack/echo/v4"
)

func (t *page) Produce(c echo.Context) (err error) {
	data, err := t.PageModel()
	if err != nil {
		return t.ServerError(c, err)
	}
	rp, err := repo.GetRepository()
	if err != nil {
		return t.ServerError(c, err)
	}
	dbZnak, err := rp.LockZnak()
	if err != nil {
		return t.ServerError(c, err)
	}
	defer rp.UnlockZnak(dbZnak)

	dt := data.Date.Format(t.Options().Layouts.TimeLayoutDay)
	data.Table, err = dbZnak.DayUtilisation(dt)
	if err != nil {
		return t.ServerError(c, err)
	}
	data.Reports = make([]*PrdReport, 0)
	if len(data.Table) > 1 {
		data.MapTable = make(map[string]map[string]int)
		for _, day := range data.Table {
			if _, exist := data.MapTable[day.ProductionDate]; !exist {
				data.MapTable[day.ProductionDate] = make(map[string]int)
			}
			if quantity, exist := data.MapTable[day.ProductionDate][day.ProductAlcCode]; exist {
				data.MapTable[day.ProductionDate][day.ProductAlcCode] = quantity + day.Quantity
			} else {
				data.MapTable[day.ProductionDate][day.ProductAlcCode] = day.Quantity
			}
		}
		for day, dayutil := range data.MapTable {
			prdReportDay, err := NewPrdReport(t)
			if err != nil {
				return t.ServerError(c, fmt.Errorf("new production report error %w", err))
			}
			prdReportDay.Report.DocProducedDate = day
			err = prdReportDay.scanDayUtilisation(dayutil)
			if err != nil {
				return t.ServerError(c, fmt.Errorf("add product in report error %w", err))
			}
			data.Reports = append(data.Reports, prdReportDay)
			err = prdReportDay.WriteBD()
			if err != nil {
				return t.ServerError(c, fmt.Errorf("new production report error %w", err))
			}
		}
	}
	// if err := c.Render(http.StatusOK, t.Name(), t.RenderPageModel("produce", data)); err != nil {
	// 	return t.ServerError(c, err)
	// }
	t.SetFlush("Отчеты созданы", "info")
	return c.NoContent(204)
}
