package summultiples

func SumMultiples(limit int, factors ...int) (sum int) {
	seen := make(map[int]bool)
	for _, factor := range factors {
		for i := factor; i < limit; i += factor {
			if !seen[i] {
				seen[i] = true
				sum += i
			}
		}
	}
	return sum
}
