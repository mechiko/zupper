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
