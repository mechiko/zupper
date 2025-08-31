package spaserver

import (
	"net/http"
)

// маршрутизация приложения
func (s *Server) Routes() http.Handler {
	s.loadViews()
	s.loadUsecaseRoutes()
	s.server.GET("/page", s.Page) // переход/загрузка на текущую страницу
	s.server.GET("/sse", s.Sse)
	return s.server
}
