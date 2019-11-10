package lc_69

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBinary(t *testing.T) {
	cases := []struct {
		given    int
		expected int
	}{
		{4, 2},
		{8, 2},
		{123456789, 11111},
	}

	for _, c := range cases {
		actual := mySqrt(c.given)
		assert.Equal(t, c.expected, actual, "given %v", c.given)
	}
}
