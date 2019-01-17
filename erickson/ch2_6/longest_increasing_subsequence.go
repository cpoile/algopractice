package ch2_6

import "math"

var arr []int

// LIS algorithm using simple backtracking. Time complexity is O(2^n) because
// it tries every combination of numbers.
func LIS(ints []int) int {
	arr = append([]int{math.MinInt64}, ints...)
	//return recLIS1(math.MinInt64, ints)
	return recLIS(0, 1)
}

func recLIS(i, j int) int {
	if j == len(arr) {
		return 0
	}
	if arr[i] >= arr[j] {
		return recLIS(i, j+1)
	} else {
		skip := recLIS(i, j+1)
		take := 1 + recLIS(j, j+1)
		return max(skip, take)
	}
}

// recursive LIS using go's slices instead of indexes
func recLIS1(prev int, ints []int) int {
	if len(ints) == 0 {
		return 0
	}
	if ints[0] <= prev {
		return recLIS1(prev, ints[1:])
	} else {
		skip := recLIS1(prev, ints[1:])
		take := 1 + recLIS1(ints[0], ints[1:])
		return max(skip, take)
	}
}

// LIS algorithm using backtracking, this time looking for the first eligible
// next increasing number. Speeds up the simple backtracking ~3x
func LISFirst(ints []int) int {
	arr = append([]int{math.MinInt64}, ints...)
	//return recLISFirst1(math.MinInt64, ints)
	return recLISFirst(0, 1)
}

func recLISFirst(i, j int) int {
	var best int
	for k := j; k < len(arr); k++ {
		if arr[i] < arr[k] {
			curBest := 1 + recLISFirst(k, k+1)
			best = max(best, curBest)
		}
	}
	return best
}

// using Go's slices instead of indexes
func recLISFirst1(prev int, ints []int) int {
	var best int
	for i, v := range ints {
		if v > prev {
			curBest := 1 + recLISFirst1(v, ints[i+1:])
			best = max(best, curBest)
		}
	}
	return best
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
