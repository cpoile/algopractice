package lc_27

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	end := len(nums)
	for i := 0; i < end; i++ {
		if nums[i] == val {
			nums[i] = nums[end-1]
			end--
			i--
		}
	}
	return end
}
