package lc_35

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	cases := []struct {
		given    []int
		target   int
		expected int
	}{
		{[]int{1, 3, 5, 6}, 4, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 7, 4},
		{[]int{1, 5, 6}, 7, 3},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1, 3, 6}, 0, 0},
		{[]int{}, 10, 0},
	}

	for _, c := range cases {
		actual := searchInsert(c.given, c.target)
		assert.Equal(t, c.expected, actual, "given %v and %d", c.given, c.target)
	}
}
