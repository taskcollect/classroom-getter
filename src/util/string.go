package util

import "strconv"

func StringToEscBuf(s string) []byte {
	return []byte(strconv.Quote(s))
}
