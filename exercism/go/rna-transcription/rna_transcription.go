package strand

import (
	"fmt"
	"strings"
)

var dToR = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	var s strings.Builder
	for _, c := range dna {
		fmt.Fprint(&s, string(dToR[c]))
	}
	return s.String()
}
