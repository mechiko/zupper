package dbs

import (
	"path/filepath"
	"regexp"
	"strings"
	"zupper/domain"
	"zupper/utility"
)

const modError = "repo:dbs"

type Dbs struct {
	domain.Apper
	// defaultDriver  string
	self           *DbInfo
	a3             *DbInfo
	znak           *DbInfo
	config         *DbInfo
	configFileName string // config.db алкохелпа
	errors         []string
}

// dbPath для своей БД
func New(apper domain.Apper, configFileName string, dbPath string) (d *Dbs) {
	var err error
	d = &Dbs{
		Apper:          apper,
		configFileName: configFileName,
		errors:         make([]string, 0),
		// defaultDriver:  logcfg.Config().Configuration().Application.DbType,
	}
	defer func() {
		if r := recover(); r != nil {
			apper.Logger().Errorf("%s %v", modError, r)
		}
	}()

	fsrarid := findA3Name()
	file4z := ""
	dbType := "sqlite"
	// по возможности получаем имена в config.db
	if dbi := d.infoDatabaseByKey("config"); dbi != nil {
		d.config = NewConfig(dbi)
		if d.config.Exists {
			// если есть база config.db то пытаемся найти настройки
			file4z = d.fromConfig("oms_id")
			dbType = strings.ToLower(d.fromConfig("db_type"))
		}
	}
	if dbi := d.infoDatabaseByKey("alcohelp3"); dbi != nil {
		if d.a3, err = NewA3(dbType, fsrarid, dbi); err != nil {
			d.Logger().Errorf("%s %s", modError, err.Error())
		}
	}
	if dbi := d.infoDatabaseByKey("trueznak"); dbi != nil {
		if d.znak, err = New4z(dbType, file4z, dbi); err != nil {
			d.Logger().Errorf("%s %s", modError, err.Error())
		}
	}
	if dbi := d.infoDatabaseByKey("selfdb"); dbi != nil {
		if d.self, err = NewSelf("", dbPath, dbi); err != nil {
			d.Logger().Errorf("%s %s", modError, err.Error())
		}
	}
	return d
}

func (d *Dbs) Self() *DbInfo {
	return d.self
}

func (d *Dbs) Znak() *DbInfo {
	return d.znak
}

func (d *Dbs) A3() *DbInfo {
	return d.a3
}

func (d *Dbs) ConfigInfo() *DbInfo {
	return d.config
}

// ни чего не делаем TODO
func (d *Dbs) SaveConfig() (err error) {
	return nil
}

func findA3DbName() string {
	re, err := regexp.Compile(`^0[0-9]{11}\.db$`)
	if err != nil {
		return ""
	}
	if files, err := utility.FilteredSearchOfDirectoryTree(re, ""); err != nil {
		return ""
	} else {
		if len(files) == 0 {
			return ""
		}
		return files[0]
	}
}

func findA3Name() string {
	findName := findA3DbName()
	if findName == "" {
		return ""
	}
	_, file := filepath.Split(findName)
	before := file[0 : len(file)-len(filepath.Ext(file))]
	// before, _ := strings.CutSuffix(file, filepath.Ext(file))
	return before
}

// key low case string
// return nil id not found
func (d *Dbs) infoDatabaseByKey(key string) *DbInfo {
	switch key {
	case "alcohelp3":
		if found, _ := utility.FindStringInJsonTags(d.Options(), key); found {
			return &DbInfo{
				Driver:     d.Options().AlcoHelp3.Driver,
				File:       d.Options().AlcoHelp3.File,
				Name:       d.Options().AlcoHelp3.DbName,
				User:       d.Options().AlcoHelp3.User,
				Pass:       d.Options().AlcoHelp3.Pass,
				Host:       d.Options().AlcoHelp3.Host,
				Port:       d.Options().AlcoHelp3.Port,
				Connection: d.Options().AlcoHelp3.Connection,
			}
		}
	case "config":
		if found, _ := utility.FindStringInJsonTags(d.Options(), key); found {
			return &DbInfo{
				Driver: d.Options().Config.Driver,
				File:   d.Options().Config.File,
				Name:   d.Options().Config.DbName,
				User:   d.Options().Config.User,
				Pass:   d.Options().Config.Pass,
				Host:   d.Options().Config.Host,
				Port:   d.Options().Config.Port,
			}
		}
	case "trueznak":
		if found, _ := utility.FindStringInJsonTags(d.Options(), key); found {
			return &DbInfo{
				Driver: d.Options().TrueZnak.Driver,
				File:   d.Options().TrueZnak.File,
				Name:   d.Options().TrueZnak.DbName,
				User:   d.Options().TrueZnak.User,
				Pass:   d.Options().TrueZnak.Pass,
				Host:   d.Options().TrueZnak.Host,
				Port:   d.Options().TrueZnak.Port,
			}
		}
	case "selfdb":
		if found, _ := utility.FindStringInJsonTags(d.Options(), key); found {
			return &DbInfo{
				Driver: d.Options().SelfDB.Driver,
				File:   d.Options().SelfDB.File,
				Name:   d.Options().SelfDB.DbName,
				User:   d.Options().SelfDB.User,
				Pass:   d.Options().SelfDB.Pass,
				Host:   d.Options().SelfDB.Host,
				Port:   d.Options().SelfDB.Port,
			}
		}
	}
	return nil
}

func (d *Dbs) Errors() []string {
	return d.errors
}

func (d *Dbs) AddError(e string) {
	d.errors = append(d.errors, e)
}
