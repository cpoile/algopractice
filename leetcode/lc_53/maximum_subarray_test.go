package lc_53

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	cases := []struct {
		given    []int
		expected int
	}{
		{[]int{-1, -1, -2, -2}, -1},
		{[]int{}, 0},
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1, 1}, 2},
		{[]int{1, 2, 1}, 4},
		{[]int{1, -1, 1, 1}, 2},
		{[]int{-2, -1, 3, 1}, 4},
	}

	for _, c := range cases {
		actual := maxSubArray(c.given)
		assert.Equal(t, c.expected, actual, "given %v", c.given)
	}
}
