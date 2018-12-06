package pangram

import "unicode"

func IsPangram(s string) bool {
	var letterSet = make(map[rune]bool)
	for _, r := range "abcdefghijklmnopqrstuvwxyz" {
		letterSet[r] = true
	}
	for _, r := range s {
		delete(letterSet, unicode.ToLower(r))
	}
	return len(letterSet) == 0
}
