package lc_67

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBinary(t *testing.T) {
	cases := []struct {
		given, given2 string
		expected      string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
		{"1111", "1", "10000"},
		{"0", "1", "1"},
		{"0", "0", "0"},
		{"1", "0", "1"},
	}

	for _, c := range cases {
		actual := addBinary(c.given, c.given2)
		assert.Equal(t, c.expected, actual, "given %v", c.given)
	}
}
