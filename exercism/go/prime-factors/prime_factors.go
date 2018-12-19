package prime

func Factors(input int64) []int64 {
	factors := make([]int64, 0)
	for i := int64(2); i <= input; {
		if input%i == 0 {
			factors = append(factors, i)
			input /= i
		} else {
			i++
		}
	}
	return factors
}
