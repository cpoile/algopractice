package lc_7

func reverse(x int) int {
	var ret int
	for x != 0 {
		ret = ret*10 + x%10
		x /= 10
	}

	if -1<<31 < ret && ret < 1<<31 {
		return ret
	}
	return 0
}
