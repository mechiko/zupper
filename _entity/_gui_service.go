package entity

import (
	"github.com/mechiko/walk"
	// "github.com/mechiko/alcogo4lite/gui/types"
	"golang.org/x/net/context"
)

type GuiService interface {
	NewMainWindow() *walk.MainWindow
	GetMainWindow() *walk.MainWindow
	// GetMPMW() MPMWInterface
	SetShutdown(f func())
	SetDefer(f func())
	SetTitle(t string)
	// SetTreeMenu(t *types.AppmenuTreeModel)
	AddAction(a *walk.Action)
	Actions() []*walk.Action
	Run(context.Context) error
	// GetGlobalInterface() models.GlobalGuiService
	// newMultiPageMainWindow() (*MultiPageMainWindow, error)
	Shutdown()
}
