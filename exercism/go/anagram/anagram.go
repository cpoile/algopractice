package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	length := len(subject)
	sortedSubj := sortString(strings.ToLower(subject))
	ret := []string{}
	for _, candidate := range candidates {
		if len(candidate) != length || strings.ToLower(candidate) == strings.ToLower(subject) {
			continue
		}
		if sortedSubj == sortString(strings.ToLower(candidate)) {
			ret = append(ret, candidate)
		}
	}
	return ret
}

func sortString(s string) string {
	t := strings.Split(s, "")
	sort.Strings(t)
	return strings.Join(t, "")
}
