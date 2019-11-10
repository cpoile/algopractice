package lc_14

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLCP(t *testing.T) {
	cases := []struct {
		given    []string
		expected string
	}{
		{[]string{}, ""},
		{[]string{""}, ""},
		{[]string{"aa", "a"}, "a"},
		{[]string{"flower"}, "flower"},
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"flower", "flow", "flight", "fudblart"}, "f"},
	}

	for _, c := range cases {
		actual := longestCommonPrefix(c.given)
		assert.Equal(t, c.expected, actual, "given %d", c.given)
	}
}
