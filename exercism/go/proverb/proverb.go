// Package proverb will produce a famous proverb with custom objects.
package proverb

import "fmt"

// Proverb will return a custom proverb with the rhyme words.
func Proverb(rhyme []string) []string {
	body := "For want of a %s the %s was lost."
	end := "And all for the want of a %s."
	sz := len(rhyme)
	var ret []string
	if sz == 0 {
		return ret
	}
	for i, word := range rhyme[:sz-1] {
		ret = append(ret, fmt.Sprintf(body, word, rhyme[i+1]))
	}
	ret = append(ret, fmt.Sprintf(end, rhyme[0]))
	return ret
}
