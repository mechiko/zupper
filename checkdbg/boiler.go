package checkdbg

import (
	"fmt"
	"zupper/repo/configdb"

	"github.com/mechiko/dbscan"
)

func (c *Checks) TestDbConfigContact() error {
	db, err := c.repo.Lock(dbscan.Config)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	if db != nil {
		// выполняются в обратном порядке
		// поэтому анлок должен быть последним выполненным
		defer c.repo.Unlock(dbscan.Config)
		defer db.Close()
	}

	dbCfg, ok := db.(*configdb.DbConfig)
	if !ok {
		return fmt.Errorf("db type wrong %T", db)
	}
	val, err := dbCfg.Key("contact_person")
	if err != nil {
		return fmt.Errorf("get key(contact_person) %w", err)
	}
	c.loger.Infof("pass TestDbConfigContact() key(contact_person) : %s", val)
	return nil
}

func (c *Checks) TestDbConfigReleaseMethod() error {
	db, err := c.repo.Lock(dbscan.Config)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer c.repo.Unlock(dbscan.Config)
	defer db.Close()

	dbCfg, ok := db.(*configdb.DbConfig)
	if !ok {
		return fmt.Errorf("db type wrong %T", db)
	}
	val, err := dbCfg.Key("release_method_type")
	if err != nil {
		return fmt.Errorf("get key(release_method_type) %w", err)
	}
	c.loger.Infof("pass TestDbConfigReleaseMethod() key(release_method_type) : %s", val)
	return nil
}

func (c *Checks) TestDbConfigContactWithoutLock() (err error) {
	info := c.repo.Info(dbscan.Config)
	if info == nil {
		return fmt.Errorf("базы config не найдено")
	}
	db, err := configdb.New(info, dbscan.Config)
	if err != nil {
		return fmt.Errorf("error open 4z db")
	}
	defer func() {
		if cerr := db.Close(); cerr != nil {
			if err != nil {
				// keep original op error and append close error
				err = fmt.Errorf("%w; close error: %v", err, cerr)
			} else {
				err = cerr
			}
		}
	}()

	val, err := db.Key("contact_person")
	if err != nil {
		return fmt.Errorf("get key(contact_person) %w", err)
	}
	c.loger.Infof("pass TestDbConfigContactWithoutLock() key(contact_person) : %s", val)
	return nil
}
