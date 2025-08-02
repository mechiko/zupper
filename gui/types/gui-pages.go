package types

import (
	"zupper/entity"

	"github.com/mechiko/walk"
)

// type PageFactoryFunc func(parent walk.Container, a IApp) (Page, error)
type PageFactoryFunc func(parent walk.Container, app IApp) (Page, error)

type Page interface {
	walk.Container
	// SetIApp(IApp)
	Parent() walk.Container
	SetParent(parent walk.Container) error
	Clear()
	// Update()
	ReductorChangeState(model entity.Model)
}

type PageConfig struct {
	Title   string
	NewPage PageFactoryFunc
	Image   string
	Class   string
}
