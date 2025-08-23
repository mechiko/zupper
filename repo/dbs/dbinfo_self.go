package dbs

import (
	"fmt"
	"path/filepath"
	"zupper/config"

	"github.com/mechiko/utility"
)

const defaultSelfDBDriver = "sqlite"

func NewSelf(dbname, dbPath string, dbi *DbInfo) (dbiOut *DbInfo, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	if dbi.Driver == "" {
		dbi.Driver = defaultSelfDBDriver
	}
	if dbi.Name == "" {
		if dbi.Name = config.Name; dbi.Name == "" {
			return dbi, fmt.Errorf("%s отсутствуют имя базы данных для Self", modError)
		}
	}
	file := filepath.Join(dbPath, fmt.Sprintf("%s.db", dbi.Name))
	if dbi.File == "" {
		dbi.File = file
	}
	// для sqlite проверяем наличие файла
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
