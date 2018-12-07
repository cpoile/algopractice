// Package luhn provides a luhn checksum validation function.
package luhn

import "unicode"

// Valid computes the Luhn checksum and returns true if it is valid.
func Valid(s string) bool {
	var count, sum int
	for i := len(s) - 1; i >= 0; i-- {
		r := rune(s[i])
		if unicode.IsSpace(r) {
			continue
		}
		if !unicode.IsDigit(r) {
			return false
		}
		cur := int(r - '0')
		if count%2 == 1 {
			cur *= 2
			if cur > 9 {
				cur -= 9
			}
		}
		sum += cur
		count++
	}
	// I put the length check here because we might get many spaces but only 0 or 1 digit.
	if count <= 1 {
		return false
	}
	return sum%10 == 0
}
