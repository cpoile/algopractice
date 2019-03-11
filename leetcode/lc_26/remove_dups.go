package lc_26

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var cur int
	for _, n := range nums[1:] {
		if n != nums[cur] {
			cur++
			nums[cur] = n
		}
	}
	return cur + 1
}
