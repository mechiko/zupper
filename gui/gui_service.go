package gui

import (
	"bytes"
	"fmt"
	"image/color"

	"zupper/domain"
	"zupper/gui/mainwindow"
	"zupper/gui/resource"
	"zupper/gui/types"
	"zupper/gui/views"
	"zupper/repo"

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
	k = 0.0
)

type guiService struct {
	domain.Apper
	repo *repo.Repository

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

func New(resourceDir string, a domain.Apper, repo *repo.Repository) (gs *guiService, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("gui New panic %v", r)
		}
	}()

	if a == nil {
		return nil, fmt.Errorf("gui:newservice argument app is nil")
	}
	if repo == nil {
		return nil, fmt.Errorf("gui:newservice argument repo is nil")
	}
	s := &guiService{
		Apper: a,
		repo:  repo,
	}
	s.niActions = make([]*walk.Action, 0)
	s.resourceDir = resourceDir
	s.Width = maxWidth
	s.Height = maxHeight
	hDC := win.GetDC(0)
	defer win.ReleaseDC(0, hDC)

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
	s.tvm = views.CreateTreeMenu(s.Apper, repo)
	if s.tvm == nil {
		return nil, fmt.Errorf("gui tree view not install")
	}
	return s, err
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

func (s *guiService) NewMainWindow() (winn *walk.MainWindow, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("gui:newmainwindow panic %v", r)
		}
	}()

	walk.Resources.SetRootDirPath(s.resourceDir)
	cfgWindow := &mainwindow.MainWindowConfig{
		Name:    "mainWindow",
		MinSize: dcl.Size{Width: int(s.Width), Height: int(s.Height)},
		MaxSize: dcl.Size{Width: int(s.Width), Height: int(s.Height)},
		OnCurrentPageChanged: func() {
			s.updateTitle(s.MainWindow.CurrentPageTitle())
		},
		Font:    dcl.Font{Family: "Verdana", PointSize: fontSizePoint},
		Visible: true,
	}

	s.MainWindow, err = mainwindow.New(s, s.repo, s.tvm)
	if err != nil {
		return nil, fmt.Errorf("gui window new %s", err)
	}

	if err = s.MainWindow.Create(cfgWindow); err != nil {
		return nil, fmt.Errorf("gui window create %s", err)
	}

	// устанавливаем иконку окна генерируем из svg
	svgIcon, err := resource.New(s).Svg(resource.SvgRequest, color.RGBA{R: 120, G: 120, B: 120, A: 255}, 64, 64)
	if err != nil {
		return nil, fmt.Errorf("mpmw:resource error %v", err)
	}
	s.MainWindow.SetIcon(svgIcon)

	// такая установка шрифта делает большую паузу при запуске программы
	// легче установить его в построителе окна
	// fontMono, err := walk.NewFont("JetBrains Mono", fontSizePoint, 0)
	// if err != nil {
	// 	return nil, fmt.Errorf("gui:walk load font %s", err.Error())
	// }
	// w.SetFont(fontMono)
	// w.AddDisposable(fontMono)

	x := int((s.ScrWidth - int32(s.MainWindow.Size().Width)) / 2)
	y := int((s.ScrHeight - int32(s.MainWindow.Size().Height)) / 2)
	s.MainWindow.SetBounds(walk.Rectangle{
		X:      x,
		Y:      y,
		Width:  s.MainWindow.Size().Width,
		Height: s.MainWindow.Size().Height,
	})

	s.updateTitle(s.MainWindow.CurrentPageTitle())

	s.MainWindow.MainWindow.Closing().Attach(s.Closing)
	return s.MainWindow.MainWindow, nil
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

func (s *guiService) Shutdown() error {
	s.MainWindow.StopTicker()
	return s.MainWindow.Close()
}
