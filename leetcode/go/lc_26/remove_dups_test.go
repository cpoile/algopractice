package lc_26

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDups(t *testing.T) {
	cases := []struct {
		given    []int
		expected []int
		length   int
	}{
		{[]int{1, 1, 2}, []int{1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}, 5},
	}

	for _, c := range cases {
		actual := removeDuplicates(c.given)
		assert.Equal(t, c.expected, c.given[:actual], "given %d", c.given)
		assert.Equal(t, c.length, actual, "given %d", c.given)
	}
}
