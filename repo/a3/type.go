package a3

import (
	"database/sql"
	"fmt"
	"zupper/domain"
	"zupper/repo/dbs"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mssql"
	"github.com/upper/db/v4/adapter/sqlite"
)

const modError = "repo:selfdb"

type DbA3 struct {
	domain.Apper
	dbSession db.Session // database session handle
}

func New(apper domain.Apper, a *dbs.DbInfo) *DbA3 {
	db := &DbA3{
		Apper: apper,
	}
	switch a.Driver {
	case "mssql":
		if a.Host == "" {
			panic(fmt.Sprintf("%s host empty", modError))
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
		db.dbSession = dbSess
		return db
	case "sqlite":
		uri := sqlite.ConnectionURL{
			Database: a.File,
			Options: map[string]string{
				"mode":          "rw",
				"_journal_mode": "DELETE",
			},
		}
		dbSess, err := sqlite.Open(uri)
		if err != nil {
			panic(fmt.Sprintf("%s %s", modError, err.Error()))
		}
		db.dbSession = dbSess
		return db
	}
	panic(fmt.Sprintf("%s unsupported or missing database driver", modError))
}

func (c *DbA3) Close() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic %v", r)
		}
	}()
	return c.dbSession.Close()
}

func (c *DbA3) DB() *sql.DB {
	if db, ok := c.dbSession.Driver().(*sql.DB); ok {
		return db
	}
	return nil
}

func (c *DbA3) Sess() db.Session {
	return c.dbSession
}

// сделано отдельно чтобы закрывать бд
func (c *DbA3) Ping() (err error) {
	defer c.dbSession.Close()
	return c.dbSession.Ping()
}
