// Package hamming provides the Hamming Distance function.
package hamming

import "fmt"

// Distance calculates the Hamming distance between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("strings are not same length")
	}
	dist := 0
	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist, nil
}
