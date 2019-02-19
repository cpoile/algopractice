package ch6

import "testing"

func TestLRS(t *testing.T) {
	var cases = []struct {
		given    string
		expected string
	}{
		{"aacaagtttacaagc", "acaag"},
		{"it was the best of times it was the", "it was the"},
	}

	for _, c := range cases {
		actual := LRS(c.given)
		if actual != c.expected {
			t.Errorf("Given: %s, expected: %s, received: %s", c.given, c.expected, actual)
		}
	}
}
