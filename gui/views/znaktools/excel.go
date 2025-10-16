package znaktools

import (
	"fmt"
	"zupper/ucexcel"
)

func (p *ZnakToolsPage) Excel(ar []string, name string) (err error) {
	excel := ucexcel.New(p, "", "", name)
	if err := excel.Open(); err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := excel.ColumnReport(ar); err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := excel.Save(name); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
