package ch3_6

import "math"

func LIS(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	arr = append([]int{math.MinInt64}, arr...)

	tab := make([][]int, len(arr))
	for k := 0; k < len(arr); k++ {
		tab[k] = make([]int, len(arr)+1)
	}

	for j := len(arr) - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if arr[i] >= arr[j] {
				tab[i][j] = tab[i][j+1]
			} else {
				tab[i][j] = max(tab[i][j+1], 1+tab[j][j+1])
			}
		}
	}
	return tab[0][1]
}

func LISFirst(ints []int) int {
	if len(ints) == 0 {
		return 0
	}
	arr := append([]int{math.MinInt64}, ints...)
	return LISFirstTable(arr)
}

func LISFirstTable(arr []int) int {
	tab := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		tab[i] = 1
	}

	for i := len(arr) - 2; i >= 0; i-- {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] && 1+tab[j] > tab[i] {
				tab[i] = 1 + tab[j]
			}
		}
	}
	return tab[0] - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
