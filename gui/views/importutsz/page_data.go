package importutsz

import (
	"time"

	"zupper/entity"
)

type pageData struct {
	Browser    string
	Pwd        string
	Output     string
	Export     string
	Limit      int
	UsePeriod  bool
	AlcoHelpDb string
	ConfigDb   string
	ZnakDb     string
	Debug      bool
	Start      time.Time
	End        time.Time
	Kvartal    string
	Periods    []string
	Year       string
	Years      []string
	//
	File      string
	FileName  string
	Input     string
	CountTtn  string
	Err       string
	Message   string
	ImportTTN entity.ImportTTN
	Fifo      bool
	Split     bool
	Reimport  bool
	// IgnoreRest bool
}

var PageData *pageData
