package encode

import (
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(s string) string {
	var count int
	var curRune rune
	ret := ""
	for _, r := range s {
		if r != curRune {
			ret += compress(curRune, count)
			curRune = r
			count = 1
		} else {
			count++
		}
	}
	ret += compress(curRune, count)
	return ret
}

func compress(curRune rune, count int) string {
	if count == 0 {
		return ""
	} else if count == 1 {
		return string(curRune)
	}
	return strconv.Itoa(count) + string(curRune)
}

func RunLengthDecode(s string) string {
	var ret string
	count := "0"
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			countInt, _ := strconv.Atoi(count)
			if countInt == 0 {
				countInt = 1 // default
			}
			ret += strings.Repeat(string(r), countInt)
			count = "0"
		} else {
			count += string(r)
		}
	}
	return ret
}
