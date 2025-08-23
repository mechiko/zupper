package dbs

import (
	"fmt"

	"github.com/mechiko/utility"
)

const defaultA3Driver = "sqlite"

func NewA3(dbType, fsrarId string, dbi *DbInfo) (dbiOut *DbInfo, err error) {
	if dbi.Driver = dbType; dbi.Driver == "" {
		dbi.Driver = defaultA3Driver
	}
	if dbi.Name == "" {
		if fsrarId == "" {
			return dbi, fmt.Errorf("%s отсутствуют имя базы данных для А3", modError)
		}
		dbi.Name = fsrarId
	} else {
		fsrarId = dbi.Name
	}
	if dbi.File == "" {
		if fsrarId == "" {
			return dbi, fmt.Errorf("%s отсутствуют имя базы данных для A3", modError)
		}
		dbi.File = fsrarId + ".db"
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
