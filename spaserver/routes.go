package spaserver

import (
	"fmt"
)

// маршрутизация приложения
func (s *Server) Routes() error {
	s.loadViews()
	err := s.loadUsecaseRoutes()
	if err != nil {
		return fmt.Errorf("routes error %w", err)
	}
	s.server.GET("/page", s.Page) // переход/загрузка на текущую страницу
	s.server.GET("/sse", s.Sse)
	return nil
}
