package lc_58

import "strings"

func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	ss := strings.Split(s, " ")
	if len(ss) > 0 {
		return len(ss[len(ss)-1])
	}
	return 0
}
