package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(input string) (Matrix, error) {
	allRows := strings.Split(input, "\n")
	rowLen := len(strings.Split(allRows[0], " "))
	var ret Matrix
	for _, r := range allRows {
		r = strings.TrimSpace(r)
		if len(strings.Split(r, " ")) != rowLen {
			return nil, errors.New("uneven rows")
		}
		row := make([]int, rowLen)
		for i, c := range strings.Split(r, " ") {
			digit, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			row[i] = digit
		}
		ret = append(ret, row)
	}
	return ret, nil
}

func (m Matrix) Rows() [][]int {
	var ret Matrix
	for _, r := range m {
		row := make([]int, len(r))
		for i, c := range r {
			row[i] = c
		}
		ret = append(ret, row)
	}
	return ret
}

func (m Matrix) Cols() [][]int {
	var ret Matrix
	for i := 0; i < len(m[0]); i++ {
		row := make([]int, len(m))
		for j := 0; j < len(m); j++ {
			row[j] = m[j][i]
		}
		ret = append(ret, row)
	}
	return ret
}

func (m Matrix) Set(r, c, val int) bool {
	rowLen := len(m)
	colLen := len(m[0])
	if r < 0 || c < 0 || r > rowLen-1 || c > colLen-1 {
		return false
	}
	m[r][c] = val
	return true
}
