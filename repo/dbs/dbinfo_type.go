package dbs

type DbInfo struct {
	Host       string
	Port       string
	User       string
	Pass       string
	File       string
	Name       string
	Driver     string
	Connection string
	Exists     bool // только для sqlite делает поиск файла
}

const defaultAddrMsSql = "localhost:1433"
