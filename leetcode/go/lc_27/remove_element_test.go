package lc_27

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	cases := []struct {
		given    []int
		val      int
		expected []int
		length   int
	}{
		{[]int{3, 2, 2, 3}, 2, []int{3, 3}, 2},
		{[]int{3, 2, 2, 2}, 2, []int{3}, 1},
		{[]int{3, 2, 3, 2}, 2, []int{3, 3}, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 1, 3, 0, 4}, 5},
	}

	for _, c := range cases {
		actual := removeElement(c.given, c.val)
		assert.ElementsMatch(t, c.expected, c.given[:actual], "given %v", c.given)
		assert.Equal(t, c.length, actual, "given %v, actual %v", c.given, c.given[:actual])
	}
}
