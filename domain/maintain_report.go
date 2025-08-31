package domain

type AdminReport struct {
	IsDoubleFormProduce        bool
	IsDoubleFormProduceRow     FormDoubleSlice
	IsDoubleFormImport         bool
	IsDoubleFormImportRow      FormDoubleSlice
	IsDoubleFormTtn            bool
	IsDoubleFormTtnRow         FormDoubleSlice
	IsDoubleForm1              bool
	IsDoubleForm1Row           FormDoubleSlice
	IsDoubleForm2              bool
	IsDoubleForm2Row           FormDoubleSlice
	AbsentForm1                AbsentItemSlice
	AbsentForm2                AbsentItemSlice
	AbsentClient               AbsentItemSlice
	IsDoubleForm2RestVolume    bool
	IsDoubleForm2RestVolumeRow FormDoubleSlice
	Errors                     AbsentItemSlice
}

type AbsentItem struct {
	Id string `db:"id"`
}

type AbsentItemSlice []*AbsentItem

type FormDouble struct {
	Id    string `db:"id"`
	Total int64  `db:"total"`
}

type FormDoubleSlice []*FormDouble
