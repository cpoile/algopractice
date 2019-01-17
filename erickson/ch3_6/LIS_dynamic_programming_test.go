package ch3_6

import (
	"testing"
)

func TestLIS(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{[]int{6, 5, 6, 7, 2, 8, 4, 9}, 5},
		{[]int{1, 4, 3, 5, 2, 4}, 3},
		{[]int{1, 5, 2, 8, 9, 1, 8, 4}, 4},
		{[]int{9, 3, 8, 4, 2, 5, 8, 3}, 4},
		{[]int{9, 8, 5, 3, 2, 1}, 1},
		{[]int{9, 4, 5, 3, 6, 10, 2, 3, 8, 5, 7, 8, 12}, 6},
		{[]int{}, 0},
		{[]int{1}, 1},
	}

	for _, c := range cases {
		actualLIS := LIS(c.input)
		if actualLIS != c.expected {
			t.Errorf("LIS: given %v, expected %v, received %v", c.input, c.expected, actualLIS)
		}
		actualLISFirst := LISFirst(c.input)
		if actualLISFirst != c.expected {
			t.Errorf("LISFirst: given %v, expected %v, received %v", c.input, c.expected, actualLISFirst)
		}
	}
}

func BenchmarkLIS(b *testing.B) {
	in := []int{1, 3, 4, 2, 30, 12, 9, 11, 5, 2, 3, 8, 3, 6, 9, 2, 5, 2, 30, 12, 9, 11, 5, 2, 3, 8, 3, 6, 9, 2, 5, 0, 1, 4, 8, 3, 6, 3, 4, 7, 1, 4, 9, 10, 23, 2, 5, 21, 4, 16, 13, 5, 2}
	for i := 0; i < b.N; i++ {
		LIS(in)
	}
}
func BenchmarkLISFirst(b *testing.B) {
	in := []int{1, 3, 4, 2, 30, 12, 9, 11, 5, 2, 3, 8, 3, 6, 9, 2, 5, 2, 30, 12, 9, 11, 5, 2, 3, 8, 3, 6, 9, 2, 5, 0, 1, 4, 8, 3, 6, 3, 4, 7, 1, 4, 9, 10, 23, 2, 5, 21, 4, 16, 13, 5, 2}
	for i := 0; i < b.N; i++ {
		LISFirst(in)
	}
}
