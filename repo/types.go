package repo

import (
	"fmt"
	"sync"

	"zupper/repo/selfdb"

	"github.com/mechiko/dbscan"
	"go.uber.org/zap"
)

const modError = "pkg:repo"

var Version int64

type singleMutex struct {
	mutex sync.Mutex
}

type Repository struct {
	logger  *zap.SugaredLogger
	dbs     *dbscan.Dbs
	dbMutex map[dbscan.DbInfoType]*singleMutex
}

// dbPath для своей БД
// func New(logcfg ILogCfg, dbPath string) (rp *Repository, err error) {
func New(logger *zap.SugaredLogger, listDbs dbscan.ListDbInfoForScan, dbPath string) (rp *Repository, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("repo panic %v", r)
		}
	}()

	dbs, err := dbscan.New(listDbs, dbPath)
	if err != nil {
		return nil, fmt.Errorf("%s dbscan error %w", modError, err)
	}
	rp = &Repository{
		logger:  logger,
		dbs:     dbs,
		dbMutex: make(map[dbscan.DbInfoType]*singleMutex),
	}
	exit := false
	for tp, dbInfo := range listDbs {
		if dbInfo == nil {
			rp.logger.Infof("%s отсутствует БД %v", modError, dbInfo)
			exit = true
		} else {
			// создаем в мапе мьютекс
			if _, ok := rp.dbMutex[tp]; !ok {
				rp.dbMutex[tp] = &singleMutex{}
			}
		}
	}
	if exit {
		return nil, fmt.Errorf("%s не все бд найдены", modError)
	}
	if di := rp.dbs.Info(dbscan.Other); di != nil {
		// миграция для Self
		if err := rp.prepareSelf(); err != nil {
			return nil, fmt.Errorf("%s ошибка миграции self %w", modError, err)
		}
	}
	return rp, nil
}

// func (r *Repository) Dbs() *dbscan.Dbs {
// 	return r.dbs
// }

// after Self() must be SelfClose() or deadlock
func (r *Repository) SelfLock() *selfdb.DbSelf {
	mu, ok := r.dbMutex[dbscan.Other]
	if ok {
		mu.mutex.Lock()
	} else {
		return nil
	}
	info := r.dbs.Info(dbscan.Other)
	if info != nil {
		return selfdb.New(r.logger, r.dbs.Info(dbscan.Other))
	}
	return nil
}

func (r *Repository) SelfUnlock() error {
	mu, ok := r.dbMutex[dbscan.Other]
	if ok {
		mu.mutex.Unlock()
	} else {
		return fmt.Errorf("%s close not present mutex %v", modError, dbscan.Other)
	}
	return nil
}

// возвращаем DbInfo или nil
func (r *Repository) Info(t dbscan.DbInfoType) *dbscan.DbInfo {
	if di := r.dbs.Info(t); di != nil {
		return di
	}
	return nil
}
