package series

import "fmt"

func All(n int, s string) []string {
	var ret []string
	for i := 0; i+n <= len(s); i++ {
		ret = append(ret, s[i:i+n])
	}
	return ret
}

func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		s = fmt.Sprintf("%s%*s", s, n-len(s), " ")
		// or: n = len(s)
	}
	return s[0:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return first, false
	}
	return s[0:n], true
}
