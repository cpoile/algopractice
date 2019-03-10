package lc_11

// brute force: n^2
func maxAreaBF(height []int) int {
	maxA := -1 << 31
	for i, s := range height {
		for j := i + 1; j < len(height); j++ {
			commonH := min(s, height[j])
			maxA = max(maxA, commonH*(j-i))
		}
	}

	return maxA
}

// outside-in
func maxArea(height []int) int {
	var maxA int
	for i, j := 0, len(height)-1; i < j; {
		common := min(height[i], height[j])
		maxA = max(maxA, common*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxA
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max3(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > a && b > c {
		return b
	}
	return c

}
