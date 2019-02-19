package ch6

import "sort"

var suffixArr = make([]string, 0)

func LRS(s string) string {
	for i := 0; i < len(s); i++ {
		suffixArr = append(suffixArr, s[i:])
	}
	sort.Strings(suffixArr)
	var longestSuf string
	for i := 1; i < len(suffixArr); i++ {
		suf := longestCommonPrefix(i-1, i)
		if len(suf) > len(longestSuf) {
			longestSuf = suf
		}
	}
	return longestSuf
}

func longestCommonPrefix(a, b int) string {
	var i int
	for i = 0; i < min(len(suffixArr[a]), len(suffixArr[b])); i++ {
		if suffixArr[a][i] != suffixArr[b][i] {
			break
		}
	}
	return suffixArr[a][:i]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
