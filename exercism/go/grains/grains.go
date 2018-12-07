// Package grains provides functions to tally the number of grains on each square
// if you were to place one grain on the first square of a chessboard, and then
// double the number of grains on each subsequent square of the board.
package grains

import "fmt"

// Square returns the number of grains of wheat that would be on this square.
func Square(square int) (uint64, error) {
	if square < 1 || square > 64 {
		return 0, fmt.Errorf("square out of range")
	}
	if square == 1 {
		return 1, nil
	}
	return 2 << uint(square-2), nil
}

// Total implements a sum of all the quantities on the chessboard positions.
func Total() uint64 {
	var sum uint64
	for i := 1; i <= 64; i++ {
		res, _ := Square(i)
		sum += res
	}
	return sum
}
