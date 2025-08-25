package znakdb

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

type DbZnak struct {
	domain.Apper
	dbSession db.Session // открытый хэндл тут
}

func New(apper domain.Apper, a *dbs.DbInfo) *DbZnak {
	db := &DbZnak{
		Apper: apper,
	}
	if a.Host == "" {
		a.Host = "localhost:1433"
	}
	switch a.Driver {
	case "mssql":
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
				"mode": "rw",
				// "_journal_mode": "WAL",
			},
		}
		dbSess, err := sqlite.Open(uri)
		if err != nil {
			panic(fmt.Sprintf("%s %s", modError, err.Error()))
		}
		db.dbSession = dbSess
		return db
	}
	panic(fmt.Sprintf("%s не указан драйвер", modError))
}

func (c *DbZnak) Close() (err error) {
	if c.dbSession == nil {
		return nil
	}
	return c.dbSession.Close()
}

func (c *DbZnak) DB() *sql.DB {
	return c.dbSession.Driver().(*sql.DB)
}

func (c *DbZnak) Sess() db.Session {
	return c.dbSession
}

// сделано отдельно чтобы закрывать бд
func (z *DbZnak) Ping() (err error) {
	defer z.dbSession.Close()
	return z.dbSession.Ping()
}
