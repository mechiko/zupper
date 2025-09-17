package domain

// table ap_egais
type ApEgais struct {
	Id                  int64  `db:"id,omitempty"`
	IdRequests          int    `db:"id_requests"`
	ProductFullName     string `db:"product_full_name"`
	ProductCapacity     string `db:"product_capacity"`
	ProductAlcVolume    string `db:"product_alc_volume"`
	ProductAlcCode      string `db:"product_alc_code"`
	ProductCode         string `db:"product_code"`
	ProductUnitType     string `db:"product_unit_type"`
	ProducerType        string `db:"producer_type"`
	ProducerClientRegId string `db:"producer_client_reg_id"`
	ProducerInn         string `db:"producer_inn"`
	ProducerKpp         string `db:"producer_kpp"`
	ProducerFullName    string `db:"producer_full_name"`
	ProducerShortName   string `db:"producer_short_name"`
	ProducerCountryCode string `db:"producer_country_code"`
	ProducerRegionCode  string `db:"producer_region_code"`
	ProducerDescription string `db:"producer_description"`
}

// production_reports
type ProductionReport struct {
	Id                  int64  `db:"id,omitempty"`
	CreateDate          string `db:"create_date"`
	DocIdentity         string `db:"doc_identity"`
	DocType             string `db:"doc_type"`
	DocNumber           string `db:"doc_number"`
	DocDate             string `db:"doc_date"`
	DocProducedDate     string `db:"doc_produced_date"`
	DocComment          string `db:"doc_comment"`
	ProducerType        string `db:"producer_type"`
	ProducerClientRegId string `db:"producer_client_reg_id"`
	ProducerInn         string `db:"producer_inn"`
	ProducerKpp         string `db:"producer_kpp"`
	ProducerFullName    string `db:"producer_full_name"`
	ProducerShortName   string `db:"producer_short_name"`
	ProducerCountryCode string `db:"producer_country_code"`
	ProducerRegionCode  string `db:"producer_region_code"`
	ProducerDescription string `db:"producer_description"`
	Version             string `db:"version"`
	State               string `db:"state"`
	Status              string `db:"status"`
	ReplyId             string `db:"reply_id"`
	Archive             int    `db:"archive"`
	Xml                 string `db:"xml"`
}

// production_products
type ProductionProduct struct {
	Id                  int64  `db:"id,omitempty"`
	IdProductionReports int64  `db:"id_production_reports"`
	ProductFullName     string `db:"product_full_name"`
	ProductCapacity     string `db:"product_capacity"`
	ProductAlcVolume    string `db:"product_alc_volume"`
	ProductAlcVolumeMin string `db:"product_alc_volume_min"`
	ProductAlcVolumeMax string `db:"product_alc_volume_max"`
	ProductAlcCode      string `db:"product_alc_code"`
	ProductCode         string `db:"product_code"`
	ProductUnitType     string `db:"product_unit_type"`
	ProductIdentity     string `db:"product_identity"`
	ProductQuantity     string `db:"product_quantity"`
	ProductParty        string `db:"product_party"`
	ProductComment      string `db:"product_comment"`
	ProducerType        string `db:"producer_type"`
	ProducerClientRegId string `db:"producer_client_reg_id"`
	ProducerInn         string `db:"producer_inn"`
	ProducerKpp         string `db:"producer_kpp"`
	ProducerFullName    string `db:"producer_full_name"`
	ProducerShortName   string `db:"producer_short_name"`
	ProducerCountryCode string `db:"producer_country_code"`
	ProducerRegionCode  string `db:"producer_region_code"`
	ProducerDescription string `db:"producer_description"`
	AsiiuQuantityDal    string `db:"asiiu_quantity_dal"`
	AsiiuQuantity       string `db:"asiiu_quantity"`
}

// partners_egais
type PartnerEgais struct {
	Id                int64  `db:"id"`
	IdRequests        int    `db:"id_requests"`
	ClientType        string `db:"client_type"`
	ClientRegId       string `db:"client_reg_id"`
	ClientInn         string `db:"client_inn"`
	ClientKpp         string `db:"client_kpp"`
	ClientFullName    string `db:"client_full_name"`
	ClientShortName   string `db:"client_short_name"`
	ClientCountryCode string `db:"client_country_code"`
	ClientRegionCode  string `db:"client_region_code"`
	ClientDescription string `db:"client_description"`
	ClientState       string `db:"client_state"`
	ClientWbVersion   string `db:"client_wb_version"`
	ClientLicense     string `db:"client_license"`
}
