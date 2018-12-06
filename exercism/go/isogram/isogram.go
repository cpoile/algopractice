// Package isogram provides a function to test if a word is an isogram.
package isogram

import "unicode"

// IsIsogram determines if a string is an isogram.
// It stops checking once a duplicate letter has been found.
func IsIsogram(s string) bool {
	counts := make(map[rune]bool)
	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}
		ru := unicode.ToUpper(r)
		if counts[ru] {
			return false
		}
		counts[ru] = true
	}
	return true
}
