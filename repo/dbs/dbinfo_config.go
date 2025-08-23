package dbs

import (
	"github.com/mechiko/utility"
)

const defaultConfigFile = "config.db"
const defaultConfigDriver = "sqlite"

// база config.db всегда sqlite3 пока
func NewConfig(dbi *DbInfo) *DbInfo {
	dbi.File = defaultConfigFile
	dbi.Driver = defaultConfigDriver
	dbi.Exists = utility.PathOrFileExists(dbi.File)
	return dbi
}
