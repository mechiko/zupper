package mainwindow

import (
	"fmt"
	"time"
	"zupper/domain"
	"zupper/gui/types"
	"zupper/repo"

	"github.com/mechiko/walk"
)

const modError = "gui:mainwindow"

type MainWindow struct {
	*walk.MainWindow
	domain.Apper
	repo *repo.Repository

	Tvm                         *types.AppmenuTreeModel
	tv                          *walk.TreeView
	pageCom                     *walk.Composite
	currentPageChangedPublisher walk.EventPublisher
	SbiScan                     *walk.StatusBarItem
	SbiFsrarId                  *walk.StatusBarItem
	SbiUtmState                 *walk.StatusBarItem
	SbiState                    *walk.StatusBarItem
	SbiLicense                  *walk.StatusBarItem
	IconRed                     *walk.Icon
	IconGreen                   *walk.Icon
	IconWatch                   *walk.Icon

	ticker        *time.Ticker
	InChangeModel chan domain.Model
}

func New(app domain.Apper, repo *repo.Repository, tree *types.AppmenuTreeModel) (*MainWindow, error) {
	if app == nil {
		return nil, fmt.Errorf("gui:new argument app is nil")
	}
	if repo == nil {
		return nil, fmt.Errorf("gui:new argument repo is nil")
	}
	w := &MainWindow{
		Apper:         app,
		repo:          repo,
		Tvm:           tree,
		InChangeModel: make(chan domain.Model, 10),
	}
	return w, nil
}
