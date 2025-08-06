package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"zupper/app"
	"zupper/checkdbg"
	"zupper/config"
	"zupper/domain/models/application"
	"zupper/gui"
	"zupper/reductor"
	"zupper/repo"
	"zupper/spaserver"
	"zupper/utility"
	"zupper/zaplog"

	"golang.org/x/sync/errgroup"
)

const modError = "main"

// var version = "0.0.0"
var fileExe string
var dir string

// если local true то папка создается локально
var local = flag.Bool("local", false, "")

func init() {
	flag.Parse()
	fileExe = os.Args[0]
	var err error
	dir, err = filepath.Abs(filepath.Dir(fileExe))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get absolute path: %v\n", err)
		os.Exit(1)
	}
	if err := os.Chdir(dir); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to change directory: %v\n", err)
		os.Exit(1)
	}
}

func errMessageExit(title string, errDescription string) {
	utility.MessageBox(title, errDescription)
	os.Exit(-1)
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	group, groupCtx := errgroup.WithContext(ctx)

	cfg, err := config.New("", !*local)
	if err != nil {
		errMessageExit("ошибка конфигурации", err.Error())
	}

	var logsOutConfig = map[string][]string{
		"logger":   {"stdout", filepath.Join(cfg.LogPath(), config.Name)},
		"echo":     {filepath.Join(cfg.LogPath(), "echo")},
		"reductor": {filepath.Join(cfg.LogPath(), "reductor")},
		"true":     {filepath.Join(cfg.LogPath(), "true")},
	}
	zl, err := zaplog.New(logsOutConfig, true)
	if err != nil {
		errMessageExit("ошибка создания логера", err.Error())
	}
	// group.Go(func() error {
	// 	return zl.Run(groupCtx)
	// })

	lg, err := zl.GetLogger("logger")
	if err != nil {
		errMessageExit("ошибка получения логера", err.Error())
	}
	loger := lg.Sugar()
	loger.Debug("zaplog started")
	loger.Infof("mode = %s", config.Mode)
	if cfg.Warning() != "" {
		loger.Infof("pkg:config warning %s", cfg.Warning())
	}

	errProcessExit := func(title string, errDescription string) {
		loger.Errorf("%s %s", title, errDescription)
		errMessageExit(title, errDescription)
	}
	// создаем приложение с опциями из конфига и логером основным
	app := app.New(cfg, loger, dir)
	// инициализируем пути необходимые приложению
	app.CreatePath()
	// создаем редуктор для хранения моделей приложения
	reductorLogger, err := zl.GetLogger("reductor")
	if err != nil {
		errProcessExit("Ошибка получения логера для редуктора", err.Error())
	}

	if err := reductor.New(reductorLogger.Sugar()); err != nil {
		errProcessExit("Ошибка создания редуктора", err.Error())
	}

	loger.Info("start repo")
	// инициализируем REPO
	// TODO изменить получение путей из конфига
	dbPath := cfg.DbPath()
	repoStart := repo.New(app, dbPath)
	if len(repoStart.Errors()) > 0 {
		fullErr := strings.Join(repoStart.Errors(), "\n")
		errProcessExit("Ошибки запуска репозитория", fullErr)
	}
	app.SetRepo(repoStart)

	appModel, err := application.New(app, repoStart)
	if err != nil {
		errProcessExit("Ошибка получения логера для редуктора", err.Error())
	}
	if err := reductor.Instance().SetModel(appModel, false); err != nil {
		errProcessExit("Ошибка редуктора", err.Error())
	}
	group.Go(func() error {
		go func() {
			<-groupCtx.Done()
			repoStart.Shutdown()
		}()
		return repoStart.Run(groupCtx)
	})
	// тесты
	if err := checkdbg.NewChecks(app).Run(); err != nil {
		loger.Errorf("check error %v", err)
		cancel()
		// Wait for cleanup to complete
		group.Wait()
		errProcessExit("Check failed", err.Error())
	}

	loger.Info("start up webapp")

	port := cfg.Configuration().HostPort
	if port == "" || port == "auto" {
		if portFree, err := utility.GetFreePort(); err == nil {
			port = fmt.Sprintf("%d", portFree)
			// порт не записываем в файл конфигурации остается только в модели приложения
			app.Config().SetInConfig("hostport", port)
		}
	}
	loger.Infof("http port %s", port)

	// тут инициализируются так же модели для всех видов
	spaServerLogger, err := zl.GetLogger("echo")
	if err != nil {
		errProcessExit("Ошибка получения логера для http server", err.Error())
	}
	httpServer := spaserver.New(app, spaServerLogger, repoStart, port, true)
	loger.Infof("отладка шаблонов %v", httpServer.TemplateIsDebug())
	loger.Infof("путь шаблонов %s", httpServer.RootPathTemplates())
	// запускаем сервер эхо через него SSE работает для флэш сообщений
	// httpServer.Start()
	group.Go(func() error {
		go func() {
			// предположим, что httpServer (как и http.ListenAndServe, кстати) не умеет останавливаться по отмене
			// контекста, тогда придётся добавить обработку отмены вручную.
			// ошибка у какого то другого члена группы или он завершился принудительно
			<-groupCtx.Done()
			app.Logger().Debugf("%s получен сигнал завершения контекста группы в HTTP", modError)
			if err := httpServer.Shutdown(); err != nil {
				app.Logger().Debugf("%s stopped http server with error: %v", modError, err)
			}
		}()
		httpServer.Start()
		// по ошибке сервера возвращаем в группу код ошибки
		return <-httpServer.Notify()
	})

	// GUI
	guiService, err := gui.New("", app, repoStart)
	if err != nil {
		// обрезаем ошибку со стеком
		errStr := err.Error()
		arrErr := strings.Split(errStr, "\n")
		if len(arrErr) > 0 {
			errStr = arrErr[0]
		}
		app.Logger().Errorf("main:gui.new error:%s", errStr)
		utility.MessageBox("ошибка создания gui", errStr)
		// по ошибке уходим на завершение программы
	} else {
		if mw, err := guiService.NewMainWindow(); err != nil {
			app.Logger().Errorf("main:walk.newmainwindows error:%s", err.Error())
			utility.MessageBox("ошибка создания главного окна", err.Error())
		} else {
			// вот теперь между созданием главного окна и запуском можем что то менять
			mw.Starting().Attach(func() {
				// действия при старте главного окна
			})
			if err := guiService.Run(ctx); err != nil {
				app.Logger().Errorf("walk:run %s", err.Error())
			}
		}
	}

	app.Logger().Info("main:walk end вызываем cancel()")
	cancel()
	// ожидание завершения всех в группе
	if err := group.Wait(); err != nil {
		fmt.Printf("game over! error %s\n", err.Error())
	} else {
		fmt.Println("game over!")
	}
	// завершаем все логи
	zl.Shutdown()
}
