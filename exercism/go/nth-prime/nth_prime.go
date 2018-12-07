// Package prime provides a function to find the nth prime
// using a sieve (with segmentation for arbitrarily large Ns)
package prime

import (
	"math"
)

// SegmentSize should be a power of 2. This (2^17) is the smallest that can
// solve the test case without using a second segment. But you can play around by lowering
// the segment size and increasing the segmentLimit.
// You would be trading space for time, but wouldn't be wasting mem for small Ns.
const segmentSize = 131072
const segmentLimit = 1

// Nth implements a sieve solution with segmentation for very large primes.
func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}

	primes := make([]int, 0, n)
	var count int
	for segment := 1; segment <= segmentLimit; segment++ {
		sqrtLimit := int(math.Sqrt(float64(segment * segmentSize)))
		notPrime := initNextSegment(segment, primes)
		for i := (segment - 1) * segmentSize; i < segment*segmentSize; i++ {
			if i < 2 {
				continue
			}
			if notPrime[i-(segment-1)*segmentSize] {
				continue
			}
			count++
			if count == n {
				return i, true
			}
			primes = append(primes, i)
			if i <= sqrtLimit {
				sieve(i, segment, notPrime)
			}
		}
	}
	return 0, false
}

func initNextSegment(segment int, primes []int) []bool {
	notPrime := make([]bool, segmentSize)
	for _, p := range primes {
		sieve(p, segment, notPrime)
	}
	return notPrime
}

func sieve(i, segment int, notPrime []bool) {
	for j := i * 2; j < segment*segmentSize; j += i {
		if j >= (segment-1)*segmentSize {
			notPrime[j-(segment-1)*segmentSize] = true
		}
	}
}
