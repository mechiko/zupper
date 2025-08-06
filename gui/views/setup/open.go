package setup

import (
	"zupper/utility"

	"github.com/skratchdot/open-golang/open"
)

func (p *HomePage) OpenDir(dir string) {
	if dir == "" {
		return
	}
	if err := open.Run(dir); err != nil {
		p.Logger().Errorf("gui:view setup open dir %s %v", dir, err)
	}
}

func (p *HomePage) Open(url string) {
	if url == "" {
		return
	}
	if err := utility.OpenURL(url); err != nil {
		p.Logger().Errorf("%v", err)
	}
}
