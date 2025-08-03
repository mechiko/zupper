package types

import (
	"zupper/domain"

	"github.com/mechiko/walk"
)

// type PageFactoryFunc func(parent walk.Container, a IApp) (Page, error)
type PageFactoryFunc func(parent walk.Container, app domain.Apper) (Page, error)

type Page interface {
	walk.Container
	Parent() walk.Container
	SetParent(parent walk.Container) error
	Clear()
}

type PageConfig struct {
	Title   string
	NewPage PageFactoryFunc
	Image   string
	Class   string
}
