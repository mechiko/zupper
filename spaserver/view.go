package spaserver

import (
	"zupper/spaserver/views/footer"
	"zupper/spaserver/views/header"
	"zupper/spaserver/views/produtil"
)

// загружаем все виды
func (s *Server) loadViews() {
	// view header
	view1 := header.New(s)
	s.views[view1.Model()] = view1
	view1.InitData()
	view1.Routes()
	// view footer
	view2 := footer.New(s)
	s.views[view2.Model()] = view2
	view2.Routes()
	view2.InitData()

	view4 := produtil.New(s)
	s.views[view4.Model()] = view4
	view4.Routes()
	view4.InitData()
}
