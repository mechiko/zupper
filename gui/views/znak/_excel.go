package znak

import (
	"fmt"

	"zupper/usecase/services/ucexcel"
)

func (p *ZnakPage) Excel(ar []string, name string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic gui excel %v", r)
		}
	}()
	excel := ucexcel.New(p, "", "", name)
	if err := excel.Open(); err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := excel.UtilizationReportList(ar); err != nil {
		return fmt.Errorf("%w", err)
	}
	if err := excel.SaveSimple(); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
