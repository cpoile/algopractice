// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate returns a tech-style acronym for the given string
func Abbreviate(s string) string {
	var ret string

	words := strings.FieldsFunc(s, func(r rune) bool { return !unicode.IsLetter(r) })
	for _, str := range words {
		for _, c := range str {
			if unicode.IsLetter(c) {
				ret += string(unicode.ToUpper(c))
				break
			}
		}
	}

	return ret
}
