package pascal

func Triangle(rows int) [][]int {
	ret := make([][]int, rows)
	for i := 0; i < rows; i++ {
		ret[i] = makeRow(i)
	}
	return ret
}

func makeRow(n int) []int {
	ret := make([]int, n+1)
	ret[0] = 1
	for i := 1; i <= n; i++ {
		ret[i] = ret[i-1] * (n + 1 - i) / i
	}
	return ret
}
