package selfdb

import (
	"database/sql"
	_ "embed"
	"fmt"

	"github.com/mechiko/dbscan"
	"github.com/upper/db/v4"
	"go.uber.org/zap"
)

const modError = "repo:selfdb"

type DbSelf struct {
	logger    *zap.SugaredLogger
	dbSession db.Session // открытый хэндл тут
}

func New(logger *zap.SugaredLogger, a *dbscan.DbInfo) *DbSelf {
	db := &DbSelf{
		logger: logger,
	}
	return db
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
