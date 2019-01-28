package transpose

import "strings"

// Transpose transposes an array of strings. Huge hattip to zeimusu for fixing this solution.
func Transpose(ss []string) []string {
	ret := []string{}
	for i := 0; ; i++ {
		row := make([]byte, len(ss))
		for r, s := range ss {
			if i < len(s) {
				row[r] = s[i]
			} else {
				row[r] = '\x00'
			}
		}
		s := strings.Replace(strings.TrimRight(string(row), "\x00"), "\x00", " ", -1)
		if len(s) == 0 {
			break
		}
		ret = append(ret, s)
	}
	return ret
}
