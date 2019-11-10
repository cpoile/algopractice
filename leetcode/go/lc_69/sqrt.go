package lc_69

func mySqrt(x int) int {
	seed := float64(x)/100 + 1
	diff := 2.0
	estimate := seed
	for diff > 1 {
		newEst := .5 * (estimate + (float64(x) / estimate))
		diff = estimate - newEst
		if diff < 0 {
			diff *= -1
		}
		estimate = newEst
	}
	return int(estimate)
}

func mySqrtSlow(x int) int {
	var n int
	for n*n <= x {
		n++
	}
	return n - 1
}
