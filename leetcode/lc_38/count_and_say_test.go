package lc_38

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	cases := []struct {
		given    int
		expected string
	}{
		{4, "1211"},
		{1, "1"},
		{2, "11"},
		{3, "21"},
		{4, "1211"},
		{5, "111221"},
	}

	for _, c := range cases {
		actual := countAndSay(c.given)
		assert.Equal(t, c.expected, actual, "given %v", c.given)
	}
}
