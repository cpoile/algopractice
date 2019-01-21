package ch3_9

import "math"

var ns, cs []int

func OptimalBST(nodes []int, costs []int) int {
	ns, cs = nodes, costs
	return OptCost(0, len(nodes)-1)
}

func OptCost(i, k int) int {
	if i > k {
		return 0
	}
	best := math.MaxInt64
	for r := i; r <= k; r++ {
		cost := OptCost(i, r-1) + OptCost(r+1, k)
		best = min(best, cost)
	}

	return sumCost(i, k) + best
}

func sumCost(i, k int) int {
	var sum int
	for j := i; j <= k; j++ {
		sum += cs[j]
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
