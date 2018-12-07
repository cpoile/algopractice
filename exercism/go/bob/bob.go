// Package bob probably should have a package comment that summarizes what it's about...
package bob

import "strings"

// Hey should also probably have a comment documenting it.
func Hey(remark string) string {
	isQuestion := strings.HasSuffix(strings.TrimSpace(remark), "?")
	hasLetters := strings.ContainsAny(strings.ToLower(remark), "abcdefghijklmnopqrstuvwxyz")
	hasNumbers := strings.ContainsAny(strings.ToLower(remark), "1233456789")
	hasSomething := hasLetters || hasNumbers
	isShouting := strings.ToUpper(remark) == remark && hasLetters

	switch {
	case isShouting && isQuestion:
		return "Calm down, I know what I'm doing!"
	case isShouting:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	case hasSomething:
		return "Whatever."
	default:
		return "Fine. Be that way!"
	}
}
