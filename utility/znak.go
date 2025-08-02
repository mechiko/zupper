package utility

import "strings"

func Serial(znak string) string {
	index := strings.IndexByte(znak, '\x1D')
	if index > 0 {
		if len(znak) > 19 {
			znak = znak[19:index]
		}
	} else {
		if len(znak) > 19 {
			znak = znak[19:]
		}
	}
	return znak
}
