package palindrome

import (
	"errors"
	"math"
	"strconv"
)

type Product struct {
	Product        int      // the palindrome
	Factorizations [][2]int // list of all possible factorizations within given limits
}

func Products(min, max int) (pmin, pmax Product, err error) {
	if min > max {
		return pmin, pmax, errors.New("fmin > fmax")
	}

	// find min
	for i := min * min; i <= max*max; i++ {
		if isPalindromic(i) {
			factors := getFactors(i, min, max)
			if len(factors) > 0 {
				pmin = Product{i, factors}
				pmax = pmin
				break
			}
		}
	}
	// find max
	for i := max * max; i >= min*min; i-- {
		if isPalindromic(i) {
			factors := getFactors(i, min, max)
			if len(factors) > 0 {
				pmax = Product{i, factors}
				break
			}
		}
	}

	if len(pmin.Factorizations) == 0 {
		return pmin, pmax, errors.New("no palindromes")
	}
	return pmin, pmax, nil
}

func getFactors(n, min, max int) [][2]int {
	var ret [][2]int
	j := minint(max, int(math.Sqrt(float64(n))))
	for i := min; i <= j; i++ {
		if n%i == 0 && n/i <= max {
			ret = append(ret, [2]int{i, n / i})
		}
	}
	return ret
}

func minint(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isPalindromic(a int) bool {
	if a < 10 {
		return true
	}
	s1 := strconv.Itoa(a)
	for i, j := 0, len(s1)-1; i < j; i, j = i+1, j-1 {
		if s1[j] != s1[i] {
			return false
		}
	}
	return true
}
