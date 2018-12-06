// Package etl provides an extract-tranform-load function for Scrabble scores.
package etl

import "strings"

// Transform takes one type of Scrabble encoding and returns the new one.
func Transform(orig map[int][]string) map[string]int {
	ret := make(map[string]int)
	for score, letters := range orig {
		for _, letter := range letters {
			ret[strings.ToLower(letter)] = score
		}
	}
	return ret
}
