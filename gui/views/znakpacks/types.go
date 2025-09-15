package znakpacks

import "strconv"

var itemPerPage = func() []string {
	out := make([]string, 0, 70)
	out = append(out, "")
	for i := 1; i <= 69; i++ {
		out = append(out, strconv.Itoa(i))
	}
	return out
}()
