package spaserver

import "zupper/usecase/adminka"

func (s *Server) loadUsecaseRoutes() {
	adminka.New(s).Routes(s.server)
}
