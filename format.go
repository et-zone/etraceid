package etraceid

import (
	"strconv"
)

func format(val int64) string {
	s := strconv.FormatInt(val, 16)
	if len(s) == 1 {
		return s + "0"
	}
	return s
}
func reduction(s string) int64 {
	v, _ := strconv.ParseInt(s, 16, 64)
	return v
}
