package armstrong

import (
	"math"
)

func IsNumber(input int) bool {
	var sum int
	orig := input
	length := int(math.Log10(float64(input))) + 1
	for input > 0 {
		digit := input % 10
		sum += int(math.Pow(float64(digit), float64(length)))
		input = input / 10
	}
	return sum == orig
}
