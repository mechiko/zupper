package mainwindow

import (
	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"
)

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
