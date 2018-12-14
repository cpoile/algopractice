// Package sieve solves the prime finding problem using an bitvector called intSet
package sieve

import "math"

type intSet struct {
	words []uint64
}

func (s *intSet) has(x int) bool {
	word, bit := x/64, uint64(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *intSet) set(x int) {
	word, bit := x/64, uint64(x%64)
	for len(s.words) <= word {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func Sieve(limit int) (ret []int) {
	var notPrimes intSet

	for i := 2; i <= limit; i++ {
		if notPrimes.has(i) {
			continue
		}

		ret = append(ret, i)
		if i <= int(math.Sqrt(float64(limit))) {
			for j := i * 2; j <= limit; j += i {
				notPrimes.set(j)
			}
		}
	}
	return
}
