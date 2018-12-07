package atbash

import (
	"strings"
	"unicode"
)

// could use bytes.Buffer and b.WriteByte
// could use []byte
// should convert to lower and remove unneeded runes in the main loop
// could be improved
func Atbash(s string) string {
	var normed string
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			normed += string(r)
		}
	}

	var count int
	var ret []string
	var tmp string
	for _, r := range normed {
		count++
		if r > '0' && r < '9' {
			tmp += string(r)
		} else {
			tmp += string('a' - r + 'z')
		}
		if count%5 == 0 {
			ret = append(ret, tmp)
			tmp = ""
		}
	}
	if len(tmp) > 0 {
		ret = append(ret, tmp)
	}
	return strings.Join(ret, " ")
}
