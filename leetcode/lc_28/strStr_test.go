package lc_28

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrStr(t *testing.T) {
	cases := []struct {
		given, needle string
		expected      int
	}{
		{"hell", "ll", 2},
		{"hello", "ll", 2},
		{"aaaaaa", "bba", -1},
		{"aaaaaa", "", 0},
		{"hellodolly", "dolly", 5},
	}

	for _, c := range cases {
		actual := strStr(c.given, c.needle)
		assert.Equal(t, c.expected, actual, "given %v", c.given)
	}
}
