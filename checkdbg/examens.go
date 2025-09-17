package checkdbg

import (
	"errors"
	"fmt"
	"time"
	"zupper/domain"
	"zupper/domain/models/application"
	"zupper/reductor"
	"zupper/repo/a3"
	"zupper/repo/configdb"
	"zupper/repo/znakdb"

	"github.com/mechiko/dbscan"
	"github.com/upper/db/v4"
	"golang.org/x/sync/errgroup"
)

func (c *Checks) TestDbConfigContact() error {
	dbCfg, err := c.repo.LockConfig()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer c.repo.UnlockConfig(dbCfg)

	val, err := dbCfg.Key("contact_person")
	if err != nil {
		return fmt.Errorf("get key(contact_person) %w", err)
	}
	c.loger.Infof("pass TestDbConfigContact() key(contact_person) : %s", val)
	return nil
}

func (c *Checks) TestDbConfigReleaseMethod() error {
	dbCfg, err := c.repo.LockConfig()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer c.repo.UnlockConfig(dbCfg)

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

func (c *Checks) TestDbA3BuilderGroupMap() error {
	dbA3, err := c.repo.LockA3()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer c.repo.UnlockA3(dbA3)

	start := time.Now()
	form1 := make([]map[string]interface{}, 0)
	session := dbA3.Sess()
	query := session.SQL().
		Select("product_inform_f1_reg_id", db.Raw("COUNT(*) as cnt")).From("form1_egais").GroupBy("product_inform_f1_reg_id")
	c.loger.Infof("sql: %s", query.String())
	err = query.All(&form1)
	if err != nil {
		return err
	}
	double := make([]string, 0)
	for _, f1 := range form1 {
		count, _ := f1["cnt"].(int64)
		if count > 1 {
			double = append(double, f1["product_inform_f1_reg_id"].(string))
		}
	}
	c.loger.Infof("pass TestDbA3BuilderGroupMap() %v", time.Since(start))
	return nil
}

func (c *Checks) TestDbA3RawGroupStruct() error {
	dbA3, err := c.repo.LockA3()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer c.repo.UnlockA3(dbA3)

	start := time.Now()
	sqlQuery := `
	select 
		product_inform_f1_reg_id as id, 
		COUNT(*) as total 
	from form1_egais 
	GROUP BY product_inform_f1_reg_id  
	HAVING count(*) > 0
	;`
	rows, errRaw := dbA3.Sess().SQL().
		Query(sqlQuery)
	if errRaw != nil {
		return fmt.Errorf("session sql error %w", errRaw)
	}
	iter := dbA3.Sess().SQL().NewIterator(rows)
	defer iter.Close()
	f1 := make(domain.FormDoubleSlice, 0)
	err = iter.All(&f1)
	if err != nil {
		return fmt.Errorf("iterator error %w", err)
	}
	c.loger.Infof("pass TestDbA3RawGroupStruct() %v", time.Since(start))
	return nil
}

func (c *Checks) TestDbA3CodeApDict() error {
	info := c.repo.Info(dbscan.A3)
	if info == nil {
		return fmt.Errorf("базы A3 не найдено")
	}
	dbA3, err := a3.New(info)
	if err != nil {
		return fmt.Errorf("error open a3 db")
	}
	defer func() {
		if cerr := dbA3.Close(); cerr != nil {
			if err != nil {
				// keep original op error and append close error
				err = fmt.Errorf("%w; close error: %v", err, cerr)
			} else {
				err = cerr
			}
		}
	}()
	ap, err := dbA3.CodeApMap()
	if err != nil {
		return err
	}
	c.loger.Infof("справочник АП %d позиций", len(ap))
	return nil
}

func (c *Checks) TestDbZnakDayUtil() error {
	info := c.repo.Info(dbscan.TrueZnak)
	if info == nil {
		return fmt.Errorf("базы A3 не найдено")
	}
	dbTz, err := znakdb.New(info)
	if err != nil {
		return fmt.Errorf("error open a3 db")
	}
	defer func() {
		if cerr := dbTz.Close(); cerr != nil {
			if err != nil {
				// keep original op error and append close error
				err = fmt.Errorf("%w; close error: %v", err, cerr)
			} else {
				err = cerr
			}
		}
	}()
	ap, err := dbTz.DayUtilisation("2025.08.29")
	if err != nil {
		return err
	}
	c.loger.Infof("сегодня нанесено %d позиций", len(ap))
	return nil
}

func (c *Checks) TestDbA3Partner() error {
	dbA3, err := c.repo.LockA3()
	if err != nil {
		err = fmt.Errorf("%s repo LockA3 error %w", modError, err)
	}
	defer func() {
		if cerr := c.repo.UnlockA3(dbA3); cerr != nil {
			err = errors.Join(err, cerr)
		}
	}()
	model, err := reductor.Instance().Model(domain.Application)
	if err != nil {
		return fmt.Errorf("get reductor model domain.Application %w", err)
	}
	mdl, ok := model.(*application.Application)
	if !ok {
		return fmt.Errorf("model wrong type %T %w", model, err)
	}
	ap, err := dbA3.PartnerByFsrarId(mdl.FsrarID)
	if err != nil {
		return err
	}
	c.loger.Infof("владелец %s %s", mdl.FsrarID, ap.ClientFullName)
	return nil
}
