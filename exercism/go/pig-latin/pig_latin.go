package piglatin

import (
	"regexp"
	"strings"
)

var reVowel = regexp.MustCompile(`^(xr|yt|[aeiou])`)
var reCons = regexp.MustCompile(`^(ch|qu|squ|thr|th|sch|rh|[bcdfghjklmnpqrstvwxyz])`)

func Sentence(input string) string {
	var ret []string
	for _, s := range strings.Split(input, " ") {
		if reVowel.MatchString(s) {
			ret = append(ret, s+"ay")
		} else if idx := reCons.FindStringIndex(s); idx != nil {
			ret = append(ret, s[idx[1]:]+s[idx[0]:idx[1]]+"ay")
		}
	}
	return strings.Join(ret, " ")
}
