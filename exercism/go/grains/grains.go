// Package grains provides functions to tally the number of grains on each square
// if you were to place one grain on the first square of a chessboard, and then
// double the number of grains on each subsequent square of the board.
package grains

import (
	"fmt"
)

// Square returns the number of grains of wheat that would be on this square.
func Square(square int) (uint64, error) {
	if square < 1 || square > 64 {
		return 0, fmt.Errorf("square out of range")
	}
	return 1 << uint(square-1), nil
}

const sumOf2to64 = 1<<64 - 1

// Total implements a sum of all the quantities on the chessboard positions.
func Total() uint64 {
	return sumOf2to64
}
