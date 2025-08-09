package repo

import (
	"database/sql"
	"fmt"
	"zupper/repo/selfdb"
	"zupper/repo/selfdb/migrations"

	"github.com/pressly/goose/v3"
)

// при инициализации приложения этот метод вызывается однажды и прописывает объект доступа
// к базе данных, далее проверяет версию БД возможна ошибка и нужно выходить из приложения
func (r *Repository) prepareSelf() (err error) {
	defer func() {
		if rr := recover(); rr != nil {
			err = fmt.Errorf("repo:dbself panic %v", rr)
		}
	}()

	self := selfdb.New(r, r.dbs.Self())
	defer self.Close()

	db := self.DB()
	dialect := r.dbs.Self().Driver
	switch dialect {
	case "sqlite":
		if err := r.makeMigrationsSqlite(db); err != nil {
			return fmt.Errorf("%s %w", modError, err)
		}
	case "mssql":
		if err := r.makeMigrationsMs(self.Sess().ConnectionURL().String()); err != nil {
			return fmt.Errorf("%s %w", modError, err)
		}
	default:
		return fmt.Errorf("%s ошибка драйвера %s", modError, dialect)
	}
	// пробуем получить версию миграции
	if Version, err = goose.GetDBVersion(db); err != nil {
		return fmt.Errorf("%s %w", modError, err)
	}
	r.SaveOptions("selfdb.driver", r.dbs.Self().Driver)
	r.SaveOptions("selfdb.file", r.dbs.Self().File)
	r.SaveOptions("selfdb.dbname", r.dbs.Self().Name)
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
