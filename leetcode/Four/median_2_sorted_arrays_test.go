package Four

import "testing"

var cases = []struct {
	a, b     []int
	expected float64
}{
	{[]int{1, 3}, []int{2}, 2.0},
	{[]int{1, 2}, []int{3, 4}, 2.5},
	{[]int{3}, []int{-2, -1}, -1.0},
}

func TestFindMedianSortedArrays(t *testing.T) {
	for _, c := range cases {
		actual := findMedianSortedArrays(c.a, c.b)
		if actual != c.expected {
			t.Errorf("Given: %v, %v, received: %f, expected %f", c.a, c.b, actual, c.expected)
		}
	}
}
