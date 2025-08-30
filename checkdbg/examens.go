package checkdbg

import (
	"fmt"
	"zupper/repo/configdb"

	"github.com/mechiko/dbscan"
	"golang.org/x/sync/errgroup"
)

func (c *Checks) TestDbConfigContact() error {
	db, err := c.repo.Lock(dbscan.Config)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer c.repo.Unlock(db)

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
	defer c.repo.Unlock(db)

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
	db, err := configdb.New(info)
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

func (c *Checks) TestDbWG() error {
	g := new(errgroup.Group)
	g.Go(c.TestDbConfigContact)
	g.Go(c.TestDbConfigReleaseMethod)
	g.Go(c.TestDbConfigContactWithoutLock)
	err := g.Wait()
	if err == nil {
		c.loger.Info("Successfully examen all test paralel.")
	}
	return err
}
