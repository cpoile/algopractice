package lc_3

func lengthOfLongestSubstring(s string) int {
	var maxLen int
	var pos [128]int
	for i := 0; i < len(pos); i++ {
		pos[i] = -1
	}
	for i, j := 0, 0; j < len(s); j++ {
		i = max(i, pos[s[j]]+1)
		pos[s[j]] = j
		maxLen = max(maxLen, j-i+1)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
