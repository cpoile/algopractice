package lc_5

var in string

func longestPalindrome2(s string) string {
	in = s
	var idx [128][]int

	for i, b := range s {
		idx[b] = append(idx[b], i)
	}

	var ret string
	for i := 0; i < len(s); i++ {
		for _, j := range idx[s[i]] {
			if j-i+1 > len(ret) && isPalindromeS(i, j) {
				ret = s[i : j+1]
			}
		}
	}
	return ret
}

func isPalindromeS(start, end int) bool {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		if in[i] != in[j] {
			return false
		}
	}
	return true
}
