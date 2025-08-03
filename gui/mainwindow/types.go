package mainwindow

import (
	"zupper/domain"
	"zupper/gui/types"

	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"
)

const modError = "gui:mainwindow"

type MainWindowConfig struct {
	Name                 string
	Enabled              dcl.Property
	Visible              dcl.Property
	Font                 dcl.Font
	MinSize              dcl.Size
	MaxSize              dcl.Size
	ContextMenuItems     []dcl.MenuItem
	OnKeyDown            walk.KeyEventHandler
	OnKeyPress           walk.KeyEventHandler
	OnKeyUp              walk.KeyEventHandler
	OnMouseDown          walk.MouseEventHandler
	OnMouseMove          walk.MouseEventHandler
	OnMouseUp            walk.MouseEventHandler
	OnSizeChanged        walk.EventHandler
	OnCurrentPageChanged walk.EventHandler
	Title                string
	Size                 dcl.Size
	MenuItems            []dcl.MenuItem
	ToolBar              dcl.ToolBar
}

type MainWindow struct {
	*walk.MainWindow
	domain.Apper
	Cfg                         *MainWindowConfig
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
}

func New(app domain.Apper, tree *types.AppmenuTreeModel) *MainWindow {

	w := &MainWindow{
		Apper: app,
		Tvm:   tree,
	}

	// регистрируем в редукторе
	// app.Reductor().RegisterGui(w.ReductorUpdaterGUI)
	return w
}
