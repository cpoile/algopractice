package lc_28

// Using Boyer-Moore, just because.
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	N, M := len(haystack), len(needle)
	const R = 128
	skip := [R]int{}

	for r := 0; r < R; r++ {
		skip[r] = -1
	}
	for i := 0; i < M; i++ {
		skip[needle[i]] = i
	}

	var shift int
	for i := 0; i <= N-M; i += shift {
		var j int
		for j = M - 1; j >= 0; j-- {
			if haystack[i+j] != needle[j] {
				shift = max(1, j-skip[haystack[i+j]])
				break
			}
		}
		if j == -1 {
			return i
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
