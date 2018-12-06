package wordcount

import (
	"regexp"
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	counts := make(Frequency)
	phrase = strings.ToLower(phrase)
	for _, w := range strings.FieldsFunc(phrase,
		func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != rune('\'')
		}) {
		w = strings.TrimPrefix(w, "'")
		w = strings.TrimSuffix(w, "'")
		counts[w]++
	}
	return counts
}

func WordCount2(phrase string) Frequency {
	ret := make(Frequency)
	reg := regexp.MustCompile(`[^\pL\d']+`)
	phrase = reg.ReplaceAllLiteralString(strings.ToLower(phrase), " ")
	phrase = regexp.MustCompile(`'([\pL\d]+)'`).ReplaceAllString(phrase, "$1")
	for _, word := range strings.Fields(phrase) {
		ret[word]++
	}
	return ret
}
