package lc_20

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidParen(t *testing.T) {
	cases := []struct {
		given    string
		expected bool
	}{
		{"", true},
		{"()", true},
		{"()[]{}", true},
		{"(}", false},
		{"([)]", false},
		{"{[]}", true},
		{"}", false},
	}

	for _, c := range cases {
		actual := isValid(c.given)
		assert.Equal(t, c.expected, actual, "given %d", c.given)
	}
}
