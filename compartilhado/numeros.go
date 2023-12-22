package compartilhado

import "strconv"

func StringParaInt(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
