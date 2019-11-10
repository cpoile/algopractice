package lc_9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		given    int
		expected bool
	}{
		{1, true},
		{100, false},
		{90000000, false},
		{121, true},
		{1221, true},
		{-121, false},
		{10, false},
		{7463883647, true},
		{746383647, true},
	}

	for _, c := range cases {
		actual := isPalindrome(c.given)
		assert.Equal(t, c.expected, actual, "given %d", c.given)
	}
}
