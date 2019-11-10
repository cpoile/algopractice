package lc_9

func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	} else if x < 10 {
		return true
	}

	var rev int
	for x > rev {
		rev = rev*10 + x%10
		x /= 10
	}

	return x == rev || x == rev/10
}
