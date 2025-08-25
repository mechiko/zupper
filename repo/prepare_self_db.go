package repo

import (
	"database/sql"
	"fmt"
	"path/filepath"
	"zupper/repo/selfdb/migrations"

	"github.com/mechiko/dbscan"
	"github.com/pressly/goose/v3"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/sqlite"
)

// при инициализации приложения этот метод вызывается однажды и прописывает объект доступа
// к базе данных, далее проверяет версию БД возможна ошибка и нужно выходить из приложения
func (r *Repository) prepareSelf() (err error) {
	defer func() {
		if rr := recover(); rr != nil {
			err = fmt.Errorf("repo:dbself panic %v", rr)
		}
	}()

	selfInfo := r.dbs.Info(dbscan.Other)
	if selfInfo == nil {
		return fmt.Errorf("repo prepareSelf база не подключена")
	}
	var self db.Session
	if !selfInfo.Exists {
		uri := selfInfo.SqliteUri(filepath.Join(selfInfo.Path, selfInfo.File))
		uri.Options["mode"] = "rwc"
		self, err = sqlite.Open(uri)
	} else {
		self, err = selfInfo.Connect()
	}
	if err != nil {
		return fmt.Errorf("repo prepareSelf ошибка подключения к БД %w", err)
	}
	defer self.Close()
	db, ok := self.Driver().(*sql.DB)
	if !ok {
		return fmt.Errorf("repo prepareSelf ошибка получения *sql.DB %T", self.Driver())
	}
	dialect := r.dbs.Info(dbscan.Other).Driver
	switch dialect {
	case "sqlite":
		if err := r.makeMigrationsSqlite(db); err != nil {
			return fmt.Errorf("%s %w", modError, err)
		}
	case "mssql":
		uri := selfInfo.MssqlUri().String()
		if err := r.makeMigrationsMs(uri); err != nil {
			return fmt.Errorf("%s %w", modError, err)
		}
	default:
		return fmt.Errorf("%s ошибка драйвера %s", modError, dialect)
	}
	// пробуем получить версию миграции
	if Version, err = goose.GetDBVersion(db); err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	return nil
}

func (r *Repository) makeMigrationsSqlite(DB *sql.DB) error {
	goose.SetBaseFS(migrations.Sqlite)
	if err := goose.SetDialect("sqlite3"); err != nil {
		return err
	}
	if err := goose.Up(DB, "sqlite"); err != nil {
		return err
	}
	return nil
}

func (r *Repository) makeMigrationsMs(uri string) error {
	goose.SetBaseFS(migrations.Mssql)
	if err := goose.SetDialect("mssql"); err != nil {
		return fmt.Errorf("failed to set MSSQL dialect: %w", err)
	}
	dbGoose, err := goose.OpenDBWithDriver("mssql", uri)
	if err == nil {
		return fmt.Errorf("failed to open MSSQL connection: %w", err)
	}
	defer dbGoose.Close()
	if err := goose.Up(dbGoose, "mssql"); err != nil {
		return fmt.Errorf("failed to run MSSQL migrations: %w", err)
	}
	return nil
}
