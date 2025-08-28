package repo

import (
	"fmt"
	"zupper/domain"
	"zupper/repo/configdb"
	"zupper/repo/selfdb"
	"zupper/repo/znakdb"

	"github.com/mechiko/dbscan"
)

// if err is nil then must after Lock launch UnLock
func (r *Repository) Lock(t dbscan.DbInfoType) (domain.RepoDB, error) {
	r.logger.Infof("repo Lock %v", t)
	info := r.dbs.Info(t)
	if info == nil {
		return nil, fmt.Errorf("repo lock dbinfo is nil for %v", t)
	}
	mu, ok := r.dbMutex[t]
	if ok {
		mu.mutex.Lock()
	} else {
		return nil, fmt.Errorf("repo lock not present mutex %v", t)
	}
	switch t {
	case dbscan.Config:
		return configdb.New(info, t)
	case dbscan.TrueZnak:
		return znakdb.New(info, t)
	case dbscan.Other:
		return selfdb.New(r.logger, info, t, false)
	default:
		mu.mutex.Unlock()
		return nil, fmt.Errorf("repo lock not present type mutex %v", t)
	}
}

func (r *Repository) Unlock(t dbscan.DbInfoType) error {
	r.logger.Infof("repo UnLock %v", t)
	mu, ok := r.dbMutex[t]
	if ok {
		mu.mutex.Unlock()
	} else {
		return fmt.Errorf("%s unlock not present mutex %v", modError, dbscan.Other)
	}
	return nil
}
