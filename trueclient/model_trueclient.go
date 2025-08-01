package trueclient

import (
	"fmt"
	"net/url"
	"time"
	"zupper/domain"
)

type TrueClientModel struct {
	Title       string
	StandGIS    url.URL
	StandSUZ    url.URL
	TokenGIS    string
	TokenSUZ    string
	AuthTime    time.Time
	LayoutUTC   string
	HashKey     string
	DeviceID    string
	OmsID       string
	IsConfigDB  bool // есть ли база конфиг.дб выставляется при запуске
	UseConfigDB bool // если ли база конфиг.дб есть то копируем данные из нее для авторизации
	Errors      []string
	PingSuz     *PingSuzInfo
	Validates   map[string]string
	MyStore     map[string]string
	Test        bool
}

type PingSuzInfo struct {
	OmsId      string `json:"omsId"`
	ApiVersion string `json:"apiVersion"`
	OmsVersion string `json:"omsVersion"`
}

func (p *PingSuzInfo) String() string {
	return fmt.Sprintf("OMS ID:%s\nAPI:%s\nOMS:%s\n", p.OmsId, p.ApiVersion, p.OmsVersion)
}

func (m *TrueClientModel) Sync(cfg domain.Apper) {
	cfg.SaveOptions("trueclient.test", m.Test)
	cfg.SaveOptions("trueclient.deviceid", m.DeviceID)
	cfg.SaveOptions("trueclient.hashkey", m.HashKey)
	cfg.SaveOptions("trueclient.omsid", m.OmsID)
	cfg.SaveOptions("trueclient.useconfigdb", m.UseConfigDB)
}

// когда считываем конфиг сбрасываем токены и время авторизации
func (m *TrueClientModel) Read(cfg domain.Apper) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
			m.Errors = append(m.Errors, err.Error())
		}
	}()
	m.TokenGIS = ""
	m.TokenSUZ = ""
	// time.IsZero()
	m.AuthTime = time.Time{}
	m.Test = cfg.Options().TrueClient.Test
	m.UseConfigDB = cfg.Options().TrueClient.UseConfigDB
	m.DeviceID = cfg.Options().TrueClient.DeviceID
	m.HashKey = cfg.Options().TrueClient.HashKey
	m.OmsID = cfg.Options().TrueClient.OmsID
	m.StandGIS = url.URL{
		Scheme: "https",
		Host:   cfg.Options().TrueClient.StandGIS,
	}
	if m.StandGIS.Host == "" {
		return fmt.Errorf("invalid or missing trueclient.standgis configuration")
	}
	m.StandSUZ = url.URL{
		Scheme: "https",
		Host:   cfg.Options().TrueClient.StandSUZ,
	}
	if m.StandSUZ.Host == "" {
		return fmt.Errorf("invalid or missing trueclient.standsuz configuration")
	}
	if m.Test {
		m.StandGIS = url.URL{
			Scheme: "https",
			Host:   cfg.Options().TrueClient.TestGIS,
		}
		m.StandSUZ = url.URL{
			Scheme: "https",
			Host:   cfg.Options().TrueClient.TestSUZ,
		}
	}

	// это делаем теперь в майн.го и в виде сетап
	// if m.IsConfigDB {
	// 	r := repo.New(cfg, "")
	// 	if len(r.Errors()) == 0 {
	// 		m.OmsID = r.ConfigDB().Key("oms_id")
	// 		m.DeviceID = r.ConfigDB().Key("connection_id")
	// 		m.HashKey = r.ConfigDB().Key("certificate_thumbprint")
	// 		m.TokenSUZ = r.ConfigDB().Key("token_suz")
	// 		m.TokenGIS = r.ConfigDB().Key("token_gis_mt")
	// 	}
	// }
	return nil
}
