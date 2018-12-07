package lsproduct

import (
	"fmt"
	"unicode"
)

func LargestSeriesProduct(s string, span int) (int, error) {
	n := len(s)
	if span == 0 {
		return 1, nil
	}
	if span < 2 || span > n {
		return 0, fmt.Errorf("span is too small")
	}
	if !onlyDigits(s) {
		return 0, fmt.Errorf("cannot handle non-digits")
	}

	var max int
	digits := sToD(s)
	for i := 0; i <= n-span; i++ {
		max = maxInt(max, product(digits[i:i+span]))
	}
	return max, nil
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func product(slice []int) int {
	p := 1
	for _, k := range slice {
		p *= k
	}
	return p
}

func sToD(s string) []int {
	d := make([]int, len(s))
	for i, r := range s {
		d[i] = int(r - '0')
	}
	return d
}

func onlyDigits(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
