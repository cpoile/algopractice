package matrix

import "math"

type Pair [2]int

// Saddle finds saddle points using three passes over the matrix.
func (m *Matrix) Saddle() []Pair {
	var pairs []Pair
	rows := make([]int, len(*m))
	cols := make([]int, len((*m)[0]))

	// one pass to find the max of each row
	for r := 0; r < len(rows); r++ {
		rows[r] = math.MinInt64
		for _, c := range (*m)[r] {
			rows[r] = max(rows[r], c)
		}
	}

	// one pass to find the max of each col
	for i := 0; i < len(cols); i++ {
		cols[i] = math.MaxInt64
		for _, r := range *m {
			cols[i] = min(cols[i], r[i])
		}
	}

	// one pass to find whether this cell is a saddle point
	for r := 0; r < len(rows); r++ {
		for c := 0; c < len(cols); c++ {
			if (*m)[r][c] >= rows[r] && (*m)[r][c] <= cols[c] {
				pairs = append(pairs, Pair{r, c})
			}
		}
	}
	return pairs
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
