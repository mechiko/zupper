package domain

type DefectItem struct {
	RegID         string  `boil:"client_reg_id"`
	Short         string  `boil:"short_name"`
	DocNumber     string  `boil:"doc_number"`
	DocDate       string  `boil:"doc_date"`
	FixDate       string  `boil:"fix_date"`
	LastActType   string  `boil:"last_act_type"`
	LastActDate   string  `boil:"last_act_date"`
	ActDiffDate   string  `boil:"act_diff_date"`
	ActDiffReject string  `boil:"act_diff_reject"`
	Volume        float64 `boil:"volume"`
	VolumeDiff    float64 `boil:"diff"`
}

type DefectItemSlice []*DefectItem
