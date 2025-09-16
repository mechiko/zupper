package spaserver

import (
	"fmt"
	"zupper/usecase/adminka"
)

func (s *Server) loadUsecaseRoutes() error {
	adm, err := adminka.New(s)
	if err != nil {
		return fmt.Errorf("spaserver load usecase error %w", err)
	}
	adm.Routes(s.server)
	return nil
}
