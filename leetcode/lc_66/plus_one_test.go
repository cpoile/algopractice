package lc_66

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlusOne(t *testing.T) {
	cases := []struct {
		given    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{0}, []int{1}},
		{[]int{4, 3, 2, 9}, []int{4, 3, 3, 0}},
		{[]int{9, 9}, []int{1, 0, 0}},
	}

	for _, c := range cases {
		actual := plusOne(c.given)
		assert.Equal(t, c.expected, actual, "given %v", c.given)
	}
}
