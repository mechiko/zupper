package dbs

import (
	"fmt"
	"strings"

	"github.com/mechiko/utility"
)

const default4ZDriver = "sqlite"

func New4z(dbType, name string, dbi *DbInfo) (dbiOut *DbInfo, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
			dbi.Exists = false
		}
	}()
	if dbi.Driver = dbType; dbi.Driver == "" {
		dbi.Driver = default4ZDriver
	}
	if dbi.File = name; dbi.File == "" {
		return dbi, fmt.Errorf("%s отсутствуют имя базы данных для 4Z", modError)
	}
	if dbi.Name = name; dbi.Name == "" {
		return dbi, fmt.Errorf("%s отсутствуют имя базы данных для 4Z", modError)
	}
	if !strings.HasSuffix(dbi.File, ".db") {
		dbi.File = dbi.File + ".db"
	}
	if dbi.Driver == "sqlite" {
		dbi.Exists = utility.PathOrFileExists(dbi.File)
	}
	if dbi.Driver == "mssql" {
		if dbi.Host == "" {
			dbi.Host = "localhost"
		}
		if dbi.Port == "" {
			dbi.Port = "1433" // default MSSQL port
		}
		dbi.Host = fmt.Sprintf("%s:%s", dbi.Host, dbi.Port)
	}
	return dbi, nil
}
