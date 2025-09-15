package domain

// элемент движения алкоголя
type HistoryItem struct {
	Type           string `boil:"type"`       // тип записи ттн производство импорт акт
	DocType        string `boil:"doc_type"`   // Расход, Приход или Возврат от покупателя, Возврат Покупателю
	DocId          int    `boil:"doc_id"`     // для ссылки на источник документ ттн акт отчет импорта производства
	DocContentId   int    `boil:"content_id"` // для ссылки на источник строка документа акта отчета ттн
	DocNumber      string `boil:"doc_number"`
	DocDate        string `boil:"doc_date"`
	DocReason      string `boil:"doc_reason"` // для актов причина списания, для ТТН расхождение если есть (отгрузка с расхождением получение с расхождением)
	DocSource      string `boil:"doc_source"`
	DocRegId       string `boil:"doc_reg_id"`
	PartnerFsrarId string `boil:"partner_fsrar_id"` // фсрар ид партнера потом берем по справочнику
	FixNumber      string `boil:"fix_number"`       // номер и дата фиксации опционально может потом потребуется
	FixDate        string `boil:"fix_date"`         // номер и дата фиксации опционально может потом потребуется
	FixDateASIIU   string `boil:"fix_date_asiiu"`   // номер и дата фиксации опционально может потом потребуется
	// описание товара
	UnitType        string  `boil:"product_unit_type"` // тип упаковки
	Quantity        float64 `boil:"quantity"`          // количество по документу
	QuantityFact    float64 `boil:"quantity_fact"`     // количество фактическое если был акт расхождения
	FullName        string  `boil:"full_name"`         // полное имя продукции
	AlcVolume       float64 `boil:"alc_volume"`        // градусы
	AlcVolumeFact   float64 `boil:"alc_volume_fact"`   // фактический градус по производству по справке А
	Code            string  `boil:"code"`              // вид АП
	AlcCode         string  `boil:"alc_code"`          // код АП
	Capacity        float64 `boil:"capacity"`          // емкость тары для Unpacked старой - 10 литров (1 дал) для новой емкость кеги в литрах
	ProducerFsrarId string  `boil:"producer_fsrar_id"` // фсрар ид производителя потом берем по справочнику
	Form1           string  `boil:"form1"`             // номер справки 1
	Form2           string  `boil:"form2"`             // номер справки 2 по владельцу утм
	BottlingDate    string  `boil:"bottling_date"`     // дара розлива
	Status          string  `boil:"status"`
	// количество по строке товара
	Counts int64   `boil:"counts"` // Packed штуки
	Dal    float64 `boil:"dal"`    // Unpacked далы
	Vol    float64 `boil:"volume"` // объем далы
	BVol   float64 `boil:"bvol"`   // безводный спирт * градусы / 100
	Summ   float64 `boil:"summ"`   // сумма по строке Quanity * Price
}

type HistoryItemSlice []*HistoryItem
