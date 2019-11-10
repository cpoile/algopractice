package lc_11

import "testing"

func TestMaxArea(t *testing.T) {
	cases := []struct {
		given    []int
		expected int
	}{
		{[]int{3, 0, 3}, 6},
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{3, 0, 0, 0, 10, 10, 0, 0, 0, 0, 3}, 30},
		{[]int{3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3}, 30},
		{[]int{3, 0, 10, 10, 0, 0, 0, 0, 0, 0, 3}, 30},
		{[]int{3, 0, 0, 0, 0, 0, 0, 10, 0, 10, 3}, 30},
		{[]int{3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 3}, 33},
		{[]int{3, 0, 0, 0, 0, 10, 10, 0, 0, 0, 0, 3}, 33},
		{[]int{3, 10, 10, 0, 0, 0, 0, 0, 0, 0, 0, 3}, 33},
		{[]int{3, 0, 0, 0, 0, 0, 0, 10, 10, 0, 0, 3}, 33},
	}

	for _, c := range cases {
		actual := maxArea(c.given)
		if actual != c.expected {
			t.Errorf("Given: %v, expected: %d, got: %d", c.given, c.expected, actual)
		}
	}
}
