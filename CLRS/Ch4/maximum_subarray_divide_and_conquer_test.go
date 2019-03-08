package Ch4

import "testing"

func TestMaxSubarray(t *testing.T) {
	cases := []struct {
		given            []int
		buy, sell, total int
	}{
		//             0    1    2   3    4    5   6   7   8    9  10   11   12  13  14  15  16
		{[]int{100, 113, 110, 85, 105, 102, 86, 63, 81, 101, 94, 106, 101, 79, 94, 90, 97},
			8, 11, 43},
		{[]int{66, 65, 63, 81, 101, 94, 106},
			3, 6, 43},
		{[]int{63, 81, 101, 94, 106},
			1, 4, 43},
		{[]int{65, 63, 81, 101, 94, 106},
			2, 5, 43},
		{[]int{100, 101, 100},
			1, 1, 1},
		{[]int{102, 101, 100},
			0, 0, 0},
		{[]int{101, 100, 101},
			2, 2, 1},
	}

	for _, c := range cases {
		s, e, tot := maxSubarray(c.given)
		if c.buy != s || c.sell != e || c.total != tot {
			t.Errorf("Given %v, wanted buy: %d, sell: %d, total: %d, received: %d, %d, %d",
				c.given, c.buy, c.sell, c.total, s, e, tot)
		}
	}
}
