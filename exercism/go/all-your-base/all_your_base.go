package allyourbase

import (
	"errors"
)

func ConvertToBase(inBase int, inDigits []int, outBase int) ([]int, error) {
	if inBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if inBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	var ret []int
	var num int
	for _, digit := range inDigits {
		if digit < 0 || digit >= inBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		num = num*inBase + digit
	}
	if num == 0 {
		ret = append(ret, 0)
	}
	for num > 0 {
		ret = append([]int{num % outBase}, ret...)
		num = num / outBase
	}
	return ret, nil
}
