package repo

import (
	"errors"
	"fmt"
	"zupper/repo/a3"

	"github.com/mechiko/dbscan"
)

// if err is nil then must after Lock launch UnLock
// всегда или открывает базу и проверяет объект или возвращает ошибку
func (r *Repository) LockA3() (*a3.DbA3, error) {
	info := r.dbs.Info(dbscan.A3)
	mu, ok := r.dbMutex[dbscan.A3]
	if ok {
		mu.mutex.Lock()
	} else {
		return nil, fmt.Errorf("repo lock not present mutex %v", dbscan.A3)
	}
	db, err := a3.New(info)
	if err != nil {
		mu.mutex.Unlock()
		return nil, fmt.Errorf("repo lock open %v error %w", dbscan.A3, err)
	}
	return db, nil
}

func (r *Repository) UnlockA3(db *a3.DbA3) error {
	if db == nil {
		return fmt.Errorf("%s unlock db %v is nil", modError, dbscan.A3)
	}
	errClose := db.Close()
	mu, ok := r.dbMutex[db.InfoType()]
	if ok {
		mu.mutex.Unlock()
	} else {
		errUnlock := fmt.Errorf("%s unlock not present mutex %v", modError, db.InfoType())
		return errors.Join(errClose, errUnlock)
	}
	return errClose
}
