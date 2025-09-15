package checkdbg

import (
	"fmt"
	"time"
	"zupper/domain"
	"zupper/repo/configdb"

	"github.com/mechiko/dbscan"
	"github.com/upper/db/v4"
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

func (c *Checks) TestDbA3BuilderGroupMap() error {
	dbA3, err := c.repo.Lock(dbscan.A3)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer c.repo.Unlock(dbA3)

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

type form1 struct {
	Form1 string `db:"id"`
	Total int64  `db:"total"`
}

func (c *Checks) TestDbA3RawGroupStruct() error {
	dbA3, err := c.repo.Lock(dbscan.A3)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer c.repo.Unlock(dbA3)

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
	// f1 := form1{}
	// for iter.Next(&f1) {
	// 	if f1.Total > 1 {
	// 		double = append(double, f1.Form1)
	// 	}
	// }
	// if err := iter.Err(); err != nil {
	// 	return fmt.Errorf("iterator error %w", err)
	// }
	c.loger.Infof("pass TestDbA3RawGroupStruct() %v", time.Since(start))
	return nil
}
