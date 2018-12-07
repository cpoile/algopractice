package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(s string) string {
	var normed []rune
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			normed = append(normed, r)
		}
	}
	if len(normed) == 0 {
		return ""
	}

	col := int(math.Ceil(math.Sqrt(float64(len(normed)))))
	row := col
	if col*(row-1) >= len(normed) {
		row--
	}

	// simpler version:
	chunks := make([]string, col)
	for i, r := range normed {
		chunks[i%col] += string(r)
	}

	return strings.Join(chunks, " ")

	// Alternative, slightly faster version:
	//
	//for i := 0; i < col; i++ {
	//	for j := 0; j < row; j++ {
	//		ret = append(ret, normed[j*col+i])
	//	}
	//	ret = append(ret, ' ')
	//}
	//// get rid of the one extra space
	//ret = ret[:len(ret)-1]
	//
	//return string(ret)
}
