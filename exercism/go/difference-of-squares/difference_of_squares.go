// Package diffsquares implements SquareOfSum, SumOfSquares, and their Difference
package diffsquares

// SquareOfSum returns the square of sums of the first n natural numbers (using Gauss's trick)
func SquareOfSum(n int) int {
	return n * (n + 1) / 2 * n * (n + 1) / 2
}

// SumOfSquares returns the sum of the squares of the first n natural numbers
// (using an extension Gauss's trick)
func SumOfSquares(n int) int {
	return (n + 1) * (2*n + 1) * n / 6
}

// Difference returns the difference of SquareOfSum and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
