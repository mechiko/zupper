package a3

import (
	"fmt"

	"github.com/mechiko/dbscan"
	"github.com/upper/db/v4"
	"go.uber.org/zap"
)

const modError = "selfdb"

type DbA3 struct {
	logger    *zap.SugaredLogger
	dbSession db.Session // открытый хэндл тут
	dbInfo    *dbscan.DbInfo
	infoType  dbscan.DbInfoType
	version   int64
}

func New(info *dbscan.DbInfo) (*DbA3, error) {
	db := &DbA3{
		dbInfo:   info,
		infoType: dbscan.A3,
	}
	if info == nil {
		return nil, fmt.Errorf("%s dbinfo is nil", modError)
	}
	// открываем сесиию в этом методе если нет ошибки
	if err := db.Check(); err != nil {
		return nil, fmt.Errorf("%s error check %v", modError, err)
	}
	return db, nil
}

func (c *DbA3) Close() (err error) {
	if c.dbSession == nil {
		return nil
	}
	return c.dbSession.Close()
}

func (c *DbA3) Sess() db.Session {
	return c.dbSession
}

func (c *DbA3) Version() int64 {
	return c.version
}

func (c *DbA3) Info() dbscan.DbInfo {
	if c.dbInfo == nil {
		return dbscan.DbInfo{}
	}
	return *c.dbInfo
}

func (c *DbA3) InfoType() dbscan.DbInfoType {
	return c.infoType
}
