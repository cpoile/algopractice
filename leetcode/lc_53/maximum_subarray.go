package lc_53

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum, curSum := -1<<31, 0
	for _, n := range nums {
		curSum = max(curSum+n, n)
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}

func maxSubArrayDivideAndConquer(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return maxArray(nums, 0, len(nums)-1)
}

func maxArray(nums []int, lo, hi int) int {
	if hi == lo {
		return nums[lo]
	}

	mid := lo + (hi-lo)/2
	lh := maxArray(nums, lo, mid)
	rh := maxArray(nums, mid+1, hi)
	cr := maxCrossing(nums, lo, mid, hi)
	if lh >= rh && lh >= cr {
		return lh
	} else if rh >= lh && rh >= cr {
		return rh
	}
	return cr
}

func maxCrossing(nums []int, lo, mid, hi int) int {
	maxLh := -1 << 31
	lSum := 0
	for l := mid; l >= lo; l-- {
		lSum += nums[l]
		maxLh = max(maxLh, lSum)
	}

	maxRh := -1 << 31
	rSum := 0
	for r := mid + 1; r <= hi; r++ {
		rSum += nums[r]
		maxRh = max(maxRh, rSum)
	}
	return maxLh + maxRh
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
