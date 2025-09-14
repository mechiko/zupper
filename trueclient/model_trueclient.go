package trueclient

import (
	"fmt"
	"net/url"
	"time"
	"zupper/domain"
)

type TrueClientModel struct {
	model       domain.Model
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

var _ domain.Modeler = (*TrueClientModel)(nil)

type PingSuzInfo struct {
	OmsId      string `json:"omsId"`
	ApiVersion string `json:"apiVersion"`
	OmsVersion string `json:"omsVersion"`
}

func (p *PingSuzInfo) String() string {
	return fmt.Sprintf("OMS ID:%s\nAPI:%s\nOMS:%s\n", p.OmsId, p.ApiVersion, p.OmsVersion)
}

func (m *TrueClientModel) Sync(cfg domain.Apper) {
	if err := cfg.SetOptions("trueclient.test", m.Test); err != nil {
		cfg.Logger().Warnw("SetOptions failed", "key", "trueclient.test", "err", err)
	}
	if err := cfg.SetOptions("trueclient.deviceid", m.DeviceID); err != nil {
		cfg.Logger().Warnw("SetOptions failed", "key", "trueclient.deviceid", "err", err)
	}
	if err := cfg.SetOptions("trueclient.hashkey", m.HashKey); err != nil {
		cfg.Logger().Warnw("SetOptions failed", "key", "trueclient.hashkey", "err", err)
	}
	if err := cfg.SetOptions("trueclient.omsid", m.OmsID); err != nil {
		cfg.Logger().Warnw("SetOptions failed", "key", "trueclient.omsid", "err", err)
	}
	if err := cfg.SetOptions("trueclient.useconfigdb", m.UseConfigDB); err != nil {
		cfg.Logger().Warnw("SetOptions failed", "key", "trueclient.useconfigdb", "err", err)
	}
	if err := m.Save(cfg); err != nil {
		cfg.Logger().Errorw("SaveOptions failed", "err", err)
	}
}

func (m *TrueClientModel) Save(cfg domain.Apper) error {
	return cfg.SaveOptions()
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

func (m *TrueClientModel) Copy() (interface{}, error) {
	// shallow copy that`s why fields is simple
	dst := *m
	dst.MyStore = make(map[string]string)
	for k, v := range m.MyStore {
		dst.MyStore[k] = v
	}
	dst.Validates = make(map[string]string)
	for k, v := range m.Validates {
		dst.Validates[k] = v
	}
	dst.Errors = make([]string, len(m.Errors))
	copy(dst.Errors, m.Errors)
	if m.PingSuz != nil {
		ping := *m.PingSuz
		dst.PingSuz = &ping
	} else {
		dst.PingSuz = nil
	}
	return &dst, nil
}

func (a *TrueClientModel) Model() domain.Model {
	return a.model
}
