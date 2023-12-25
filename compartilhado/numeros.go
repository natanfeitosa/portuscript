package compartilhado

import "strconv"

func StringParaInt(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func StringParaDec(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
