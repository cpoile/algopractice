package lc_66

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; carry > 0; i-- {
		if i < 0 {
			digits = append([]int{carry}, digits...)
			break
		}
		carry = (digits[i] + 1) / 10
		digits[i] = (digits[i] + 1) % 10
	}
	return digits
}
