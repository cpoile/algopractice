package lc_5

func longestPalindrome(s string) string {
	var ans string
	var curi, curj int
	for i := 0; i < len(s); i++ {
		// expand around the centre, centre = i
		for j := 0; i-j >= 0 && i+j < len(s); j++ {
			if s[i-j] == s[i+j] {
				curi = i - j
				curj = i + j
			} else {
				break
			}
		}
		if curj-curi+1 > len(ans) {
			ans = s[curi : curj+1]
		}

		// expand around the centre, centre = i+.5
		for j, k := i, i+1; j >= 0 && k < len(s); j, k = j-1, k+1 {
			if s[j] == s[k] {
				curi = j
				curj = k
			} else {
				break
			}
		}
		if curj-curi+1 > len(ans) {
			ans = s[curi : curj+1]
		}
	}
	return ans
}
