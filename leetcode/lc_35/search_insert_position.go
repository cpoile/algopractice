package lc_35

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	lo, hi := 0, len(nums)-1
	for hi >= lo {
		mid := lo + (hi-lo)/2
		if target < nums[mid] {
			hi = mid - 1
		} else if target > nums[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return hi + 1
}
