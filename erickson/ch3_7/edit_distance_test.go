package ch3_7

import (
	"math/rand"
	"testing"
	"time"
)

func TestEditDistance(t *testing.T) {
	cases := []struct {
		s1, s2   string
		expected int
	}{
		{"algorithm", "altruistic", 6},
		{"", "", 0},
		{"l", "", 1},
		{"", "l", 1},
		{"little", "little", 0},
		{"little", "littl", 1},
		{"little", "litt", 2},
		{"little", "littler", 1},
		{"little", "littlery", 2},
		{"little", "littlery one", 6},
		{"jBeiUDzWIHw", "ggWUvnFmmt", 10},
		{"ab", "acd", 2},
	}

	for _, c := range cases {
		actual := EditDistance(c.s1, c.s2)
		if actual != c.expected {
			t.Errorf("Recursive: Given %s -> %s, expected: %v, got: %v", c.s1, c.s2, c.expected, actual)
		}
		actualDP := EditDistanceDP(c.s1, c.s2)
		if actual != c.expected {
			t.Errorf("DP: Given %s -> %s, expected: %v, got: %v", c.s1, c.s2, c.expected, actualDP)
		}
	}
}

func BenchmarkEditDistanceRecursive(t *testing.B) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Seed for Recursive: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < t.N; i++ {
		EditDistance(genRandString(rng, 10), genRandString(rng, 10))
	}
}

func BenchmarkEditDistanceDP(t *testing.B) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Seed for Recursive: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < t.N; i++ {
		EditDistanceDP(genRandString(rng, 10), genRandString(rng, 10))
	}
}

func genRandString(rng *rand.Rand, sz int) string {
	bytes := make([]byte, sz)
	for i := 0; i < sz; i++ {
		bytes[i] = byte('a' + rng.Intn(26))
	}
	return string(bytes)
}
