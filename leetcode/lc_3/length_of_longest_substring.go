package lc_3

func lengthOfLongestSubstring1(s string) int {
	var seen [127]bool
	var maxLen, shiftTo, curCount int
	var history []byte
	for i := 0; i < len(s); i = shiftTo {
		for j := i; j < len(s); j++ {
			// append no matter what
			history = append(history, s[j])
			if !seen[s[j]] {
				seen[s[j]] = true
				curCount++
			} else {
				shiftTo = j + 1
				break
			}
		}
		maxLen = max(maxLen, curCount)

		// remove seen vals until we reach the duplicate
		for history[0] != history[len(history)-1] {
			seen[history[0]] = false
			history = history[1:]
			curCount--
		}
		// get rid of duplicate
		history = history[1:]
	}
	return maxLen
}

func max2(a, b int) int {
	if a > b {
		return a
	}
	return b
}
