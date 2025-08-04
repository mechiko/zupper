package types

import (
	"zupper/domain"
	"zupper/repo"

	"github.com/mechiko/walk"
)

// type PageFactoryFunc func(parent walk.Container, a IApp) (Page, error)
type PageFactoryFunc func(parent walk.Container, app domain.Apper, repo *repo.Repository) (Page, error)

type Page interface {
	walk.Container
	Parent() walk.Container
	SetParent(parent walk.Container) error
	Clear()
	PageModel() interface{} // *PageModel указатель возвращаем
	InitData() interface{}  // *PageModel указатель возвращаем
	Update()
	SetSendFunc(f func(domain.Model))
}

type PageConfig struct {
	Title   string
	NewPage PageFactoryFunc
	Image   string
	Class   string
}
