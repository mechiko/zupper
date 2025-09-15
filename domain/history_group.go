package domain

type HistoryGroupItem struct {
	Items        HistoryItemSlice
	SubGroup     MapHistoryGroup
	SubGroupKeys []string
	Title        string
	IsOutgo      bool // true если расход
	Counts       int64
	Dal          float64
	Vol          float64
	BVol         float64
	Summ         float64
}

type MapHistoryGroup map[string]*HistoryGroupItem
