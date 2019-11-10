package lc_7

import "testing"

func TestMaxArea(t *testing.T) {
	cases := []struct {
		given    int
		expected int
	}{
		{-1234, -4321},
		{123, 321},
		{-123, -321},
		{120, 21},
		{-120, -21},
		{463847412, 214748364},
		{2147483647, 0},
		{1563847412, 0},
	}

	for _, c := range cases {
		actual := reverse(c.given)
		if actual != c.expected {
			t.Errorf("Given: %v, expected: %d, got: %d", c.given, c.expected, actual)
		}
	}
}
