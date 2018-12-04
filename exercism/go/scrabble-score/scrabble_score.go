// Package scrabble provides a simple scrabble scoring function.
package scrabble

import (
	"unicode"
)

// Score returns the Scrabble score for a given word.
func Score(s string) int {
	var total int
	for _, r := range s {
		switch unicode.ToUpper(r) {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			total++
		case 'D', 'G':
			total += 2
		case 'B', 'C', 'M', 'P':
			total += 3
		case 'F', 'H', 'V', 'W', 'Y':
			total += 4
		case 'K':
			total += 5
		case 'J', 'X':
			total += 8
		case 'Q', 'Z':
			total += 10
		}
	}
	return total
}
