package gui

import (
	"bytes"
	"fmt"
	"image/color"

	"zupper/gui/mainwindow"
	"zupper/gui/resource"
	"zupper/gui/types"
	"zupper/gui/views"

	"github.com/lxn/win"
	"github.com/mechiko/walk"
	dcl "github.com/mechiko/walk/declarative"
)

const (
	fontSizePoint       = 11
	maxWidth      int32 = 1100
	maxHeight     int32 = 700
	// коэффициент уменьшения ширины и высоты, расчитывается как максимальная ширина реального экрана деленная на этот коэф
	// k        float32 = float32(3) / float32(5)
	k        = 0.0
	modError = "gui"
)

type guiService struct {
	types.IApp
	shutdown    func()
	ondefer     func()
	resourceDir string
	title       string
	MainWindow  *mainwindow.MainWindow
	NotifyIcon  *walk.NotifyIcon
	tvm         *types.AppmenuTreeModel
	niActions   []*walk.Action
	ScrWidth    int32
	ScrHeight   int32
	Width       int32
	Height      int32
	X           int32
	Y           int32
}

// var _ entity.GuiService = &guiService{}

func New(resourceDir string, a types.IApp) entity.GuiService {
	defer func() {
		if r := recover(); r != nil {
			panic(fmt.Errorf("%s New panic %v", modError, r))
		}
	}()

	s := &guiService{
		IApp: a,
	}
	s.niActions = make([]*walk.Action, 0)
	s.resourceDir = resourceDir
	s.Width = maxWidth
	s.Height = maxHeight
	hDC := win.GetDC(0)
	defer win.ReleaseDC(0, hDC)

	// s.ScrWidth = int(win.GetDeviceCaps(hDC, win.HORZRES))
	// s.ScrHeight = int(win.GetDeviceCaps(hDC, win.VERTRES))
	s.ScrWidth = win.GetSystemMetrics(win.SM_CXSCREEN)
	s.ScrHeight = win.GetSystemMetrics(win.SM_CYSCREEN)
	if k > 0 {
		if s.ScrWidth < maxWidth {
			s.Width = s.ScrWidth - 100
		} else {
			s.Width = int32(float32(s.ScrWidth) * k)
		}
		if s.ScrHeight < maxHeight {
			s.Height = s.ScrHeight - 50
		} else {
			s.Height = int32(float32(s.ScrHeight) * k)
		}
	}
	if (s.ScrWidth - s.Width) > 0 {
		s.X = (s.ScrWidth - s.Width) / 2
	} else {
		s.X = 0
	}
	if (s.ScrHeight - s.Height) > 0 {
		s.Y = (s.ScrHeight - s.Height) / 2
	} else {
		s.Y = 0
	}
	// дерево меню инициализируем до создания главного окна
	s.tvm = views.CreateTreeMenu(s.IApp)
	return s
}

func (s *guiService) AddAction(a *walk.Action) {
	s.niActions = append(s.niActions, a)
}

func (s *guiService) Actions() []*walk.Action {
	return s.niActions
}

func (s *guiService) SetDefer(f func()) {
	s.ondefer = f
}

func (s *guiService) SetShutdown(f func()) {
	s.shutdown = f
}

func (s *guiService) NewMainWindow() *walk.MainWindow {
	var err error
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		err = fmt.Errorf("gui:mainwindow panic %v", r)
	// 	}
	// }()

	walk.Resources.SetRootDirPath(s.resourceDir)
	cfg := &mainwindow.MainWindowConfig{
		Name:    "mainWindow",
		MinSize: dcl.Size{Width: int(s.Width), Height: int(s.Height)},
		MaxSize: dcl.Size{Width: int(s.Width), Height: int(s.Height)},
		OnCurrentPageChanged: func() {
			s.updateTitle(s.MainWindow.CurrentPageTitle())
		},
		Visible: true,
	}

	w := mainwindow.New(s, s.tvm)

	if s.tvm == nil {
		panic(fmt.Errorf("%s tree view not install", modError))
	}

	w.Cfg = cfg

	if err = w.Create(); err != nil {
		panic(fmt.Errorf("%s %s", modError, err))
	}
	s.MainWindow = w

	svgIcon, err := resource.New(s).Svg(resource.SvgRequest, color.RGBA{R: 120, G: 120, B: 120, A: 255}, 64, 64)
	if err != nil {
		s.Logger().Errorf("mpmw:resource error %s", err)
	}
	w.SetIcon(svgIcon)

	fontMono, err := walk.NewFont("JetBrains Mono", fontSizePoint, 0)
	if err != nil {
		s.Logger().Errorf("gui:walk load font %s", err.Error())
	}
	w.SetFont(fontMono)
	w.AddDisposable(fontMono)

	x := int((s.ScrWidth - int32(w.Size().Width)) / 2)
	y := int((s.ScrHeight - int32(w.Size().Height)) / 2)
	w.SetBounds(walk.Rectangle{
		X:      x,
		Y:      y,
		Width:  w.Size().Width,
		Height: w.Size().Height,
	})

	s.updateTitle(s.MainWindow.CurrentPageTitle())

	w.MainWindow.Closing().Attach(s.Closing)
	// i := 1
	// s.ProgressDialog(w, func() int {
	// 	i += 5
	// 	return i
	// })

	w.SetVisible(true)
	win.SetForegroundWindow(w.Handle())
	return w.MainWindow
}

func (s *guiService) Closing(canceled *bool, reason walk.CloseReason) {
	s.Logger().Infof("gui:mainwindows Closing() canceled=%v reason=%+v", *canceled, reason)
	walk.App().Exit(0)
}

func (s *guiService) Defer() {
	if s.NotifyIcon != nil {
		s.NotifyIcon.Dispose()
	}
	if s.ondefer != nil {
		s.ondefer()
	}
}

func (s *guiService) SetTitle(t string) {
	s.title = t
}

func (s *guiService) updateTitle(prefix string) {
	var buf bytes.Buffer

	if prefix != "" {
		buf.WriteString(prefix)
		buf.WriteString(" - ")
	}

	buf.WriteString(s.title)

	s.MainWindow.SetTitle(buf.String())
}

func (s *guiService) GetCurrentPage() (*types.Page, error) {
	currentPage := s.tvm.CurrentPage()
	return &currentPage, nil
}

func (s *guiService) GetMainWindow() *walk.MainWindow {
	return s.MainWindow.MainWindow
}

func (s *guiService) Shutdown() {
	s.MainWindow.Close()
}
