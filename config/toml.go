package config

var TomlConfig = []byte(`
# This is a TOML document.
hostname = "127.0.0.1"
hostport = "8888"
output = "output"
browser = ""

[layouts]
timelayout = "2006-01-02T15:04:05-0700"
timelayoutclear = "2006.01.02 15:04:05"
timelayoutday = "2006.01.02"
timelayoututc = "2006-01-02T15:04:05"

[application]
license = "f7bc886d-bbcd-4ce9-845f-1209d87d406d"
fsrarid = ""
dbtype = "sqlite"

[trueclient]
test = false
layoututc = "2006-01-02T15:04:05"
standgis = "markirovka.crpt.ru"
standsuz = "suzgrid.crpt.ru"
testgis = "markirovka.sandbox.crptech.ru"
testsuz = "suz.sandbox.crptech.ru"
deviceid = ''
hashkey = ''
omsid = ''
useconfigdb = true

[[dbs]]
name = 'config'
dbname = ''
driver = 'sqlite'
file = ''

[[dbs]]
name = 'az'
dbname = ''
driver = 'sqlite'
file = ''

[[dbs]]
name = 'a3'
dbname = ''
driver = 'sqlite'
file = ''

`)

type Configuration struct {
	Hostname string `mapstructure:"hostname"`
	HostPort string `mapstructure:"hostport"`
	Output   string `mapstructure:"output"`
	Export   string `mapstructure:"export"`
	Browser  string `mapstructure:"browser"`

	Application AppConfiguration    `mapstructure:"application"`
	Layouts     LayoutConfiguration `mapstructure:"layouts"`
	// описатели БД рефактор
	Dbs []DatabaseConfiguration `mapstructure:"dbs"`
	// Config    DatabaseConfiguration `mapstructure:"config"`
	// AlcoHelp3 DatabaseConfiguration `mapstructure:"alcohelp3"`
	// TrueZnak  DatabaseConfiguration `mapstructure:"trueznak"`
	// SelfDB    DatabaseConfiguration `mapstructure:"selfdb"`
	// описание клиента ЧЗ
	TrueClient TrueClientConfig `mapstructure:"trueclient"`
}

type LayoutConfiguration struct {
	TimeLayout      string `mapstructure:"timelayout"`
	TimeLayoutClear string `mapstructure:"timelayoutclear"`
	TimeLayoutDay   string `mapstructure:"timelayoutday"`
	TimeLayoutUTC   string `mapstructure:"timelayoututc"`
}

type DatabaseConfiguration struct {
	Name       string `mapstructure:"name"`
	Connection string `mapstructure:"connection"`
	Driver     string `mapstructure:"driver"`
	DbName     string `mapstructure:"dbname"`
	File       string `mapstructure:"file"`
	User       string `mapstructure:"user"`
	Pass       string `mapstructure:"pass"`
	Host       string `mapstructure:"host"`
	Port       string `mapstructure:"port"`
}

type AppConfiguration struct {
	// Pwd          string `mapstructure:"pwd"`
	// Console      bool   `mapstructure:"console"`
	// Disconnected bool   `mapstructure:"disconnected"`
	Fsrarid string `mapstructure:"fsrarid"`
	// DbType       string `mapstructure:"dbtype"`
	License string `mapstructure:"license"`
	// ScanTimer    int    `mapstructure:"scantimer"`
	StartPage string `mapstructure:"startpage"`
}

type TrueClientConfig struct {
	Test        bool   `mapstructure:"test"`
	StandGIS    string `mapstructure:"standgis"`
	StandSUZ    string `mapstructure:"standsuz"`
	TestGIS     string `mapstructure:"testgis"`
	TestSUZ     string `mapstructure:"testsuz"`
	TokenGIS    string `mapstructure:"tokengis"`
	TokenSUZ    string `mapstructure:"tokensuz"`
	AuthTime    string `mapstructure:"authtime"`
	LayoutUTC   string `mapstructure:"layoututc"`
	HashKey     string `mapstructure:"hashkey"`
	DeviceID    string `mapstructure:"deviceid"`
	OmsID       string `mapstructure:"omsid"`
	UseConfigDB bool   `mapstructure:"useconfigdb"`
}
