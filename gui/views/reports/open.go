package reports

import (
	"zupper/utility"

	"github.com/skratchdot/open-golang/open"
)

func (p *ReportsPage) OpenDir(dir string) {
	if dir == "" {
		return
	}
	if err := open.Run(dir); err != nil {
		p.Logger().Errorf("gui:view setup open dir %s %v", dir, err)
	}
}

func (p *ReportsPage) Open(url string) {
	if url == "" {
		return
	}
	if err := utility.OpenURL(url); err != nil {
		p.Logger().Errorf("%v", err)
	}
}
