package ch3_9

import "testing"

func TestOptimalBST(t *testing.T) {
	cases := []struct {
		nodes    []int
		costs    []int
		expected int
	}{
		{[]int{10, 12}, []int{34, 50}, 118},
		{[]int{10, 12, 20}, []int{34, 8, 50}, 142},
	}

	for _, c := range cases {
		actual := OptimalBST(c.nodes, c.costs)
		if actual != c.expected {
			t.Errorf("given nodes %v, with costs %v, expected: %d, actual: %d", c.nodes, c.costs, c.expected, actual)
		}
	}
}
