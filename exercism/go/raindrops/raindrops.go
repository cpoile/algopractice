// Package raindrops provides the Convert function,
// which converts an int to raindrop sounds.
package raindrops

import (
	"strconv"
)

// Convert takes an int and turns it into raindrop sounds if it has the factors 3, 5, or 7.
func Convert(x int) string {
	var s string
	if x%3 == 0 {
		s += "Pling"
	}
	if x%5 == 0 {
		s += "Plang"
	}
	if x%7 == 0 {
		s += "Plong"
	}
	if len(s) == 0 {
		s = strconv.Itoa(x)
	}

	return s
}
