package selfdb

import (
	"database/sql"
	_ "embed"
	"fmt"
	"zupper/domain"
	"zupper/repo/dbs"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mssql"
	"github.com/upper/db/v4/adapter/sqlite"
)

const modError = "repo:selfdb"

type DbSelf struct {
	domain.Apper
	dbSession db.Session // открытый хэндл тут
}

func New(apper domain.Apper, a *dbs.DbInfo) *DbSelf {
	db := &DbSelf{
		Apper: apper,
	}
	switch a.Driver {
	case "mssql":
		if a.Host == "" {
			a.Host = "localhost"
		}
		if a.Port == "" {
			a.Port = "1433"
		}
		uri := mssql.ConnectionURL{
			User:     a.User,
			Password: a.Pass,
			Host:     a.Host,
			Database: a.Name,
			Options: map[string]string{
				"encrypt": "disable",
			},
		}
		dbSess, err := mssql.Open(uri)
		if err != nil {
			panic(fmt.Sprintf("%s %s", modError, err.Error()))
		}
		a.Exists = true
		db.dbSession = dbSess
		return db
	case "sqlite":
		uri := sqlite.ConnectionURL{
			Database: a.File,
			Options: map[string]string{
				"mode":          "rwc",
				"_journal_mode": "DELETE",
			},
		}
		dbSess, err := sqlite.Open(uri)
		if err != nil {
			panic(fmt.Sprintf("%s %s", modError, err.Error()))
		}
		a.Exists = true
		db.dbSession = dbSess
		return db
	}
	panic(fmt.Sprintf("%s не указан драйвер", modError))
}

func (c *DbSelf) Close() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
		}
	}()
	return c.dbSession.Close()
}

func (c *DbSelf) DB() *sql.DB {
	return c.dbSession.Driver().(*sql.DB)
}

func (c *DbSelf) Sess() db.Session {
	return c.dbSession
}

// сделано отдельно чтобы закрывать бд
func (c *DbSelf) Ping() (err error) {
	defer c.dbSession.Close()
	return c.dbSession.Ping()
}
