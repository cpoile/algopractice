package lc_58

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	cases := []struct {
		given    string
		expected int
	}{
		{"Hello World", 5},
		{"Hello World this is base", 4},
		{"baseare", 7},
		{"", 0},
		{"a ", 1},
		{"a               ", 1},
		{"               ", 0},
	}

	for _, c := range cases {
		actual := lengthOfLastWord(c.given)
		assert.Equal(t, c.expected, actual, "given %v", c.given)
	}
}
