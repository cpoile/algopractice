package lc_5

import (
	"testing"
)

var cases = []struct {
	input    string
	expected string
}{
	{"a", "a"},
	{"babad", "bab"},
	{"cbbd", "bb"},
	{"aorsitarstarstarstarienaamanaplanacanalpanamarnsoietnarosietnaroisetnoarisueaoriestnoarisetnoaiersntoiaresntoiraesntenawyuflpaioersntyuwlfpoiealrstoiewanlfpyuoanrstyunwfaopielarsy;tuonwfpyuanorstinlawfyupnatrosyutn", "amanaplanacanalpanama"},
}

func TestLongestPalindrome(t *testing.T) {
	for _, c := range cases {
		actual := longestPalindrome(c.input)
		if actual != c.expected {
			t.Errorf("given: %s, received: %s, expected: %s", c.input, actual, c.expected)
		}
	}
}

func BenchmarkLongest(t *testing.B) {
	for i := 0; i < t.N; i++ {
		for _, c := range cases {
			longestPalindrome(c.input)
		}
	}
}
