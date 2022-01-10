package util

import (
	"strings"
)

func MakePlaceholderJSONArray(length int) []byte {
	// FIXME: THIS HAS TO BE VERY UNSAFE OH GOD
	// please dont hack us <3

	// what this does, given length let's say 5
	// [0,0,0,0,0]
	// so jsonparser.Set can work with indexes, this makes placeholders
	// zeros will be literally replaced with whatever jsonparser thinks is good

	if length <= 0 {
		return []byte{'[', ']'}
	}

	arr_s := ("[" + strings.Repeat("0,", length))
	arr_s = arr_s[:len(arr_s)-1] + "]" // remove trailing "," and add ]

	return []byte(arr_s)
}
