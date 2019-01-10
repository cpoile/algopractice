package wordy

import (
	"strconv"
	"strings"
)

var acceptedOps = map[string]bool{"plus": true, "minus": true, "multiplied": true, "divided": true}

func Answer(q string) (int, bool) {
	if !strings.HasPrefix(q, "What is ") {
		return 0, false
	}
	q = q[8 : len(q)-1]
	tokens := strings.Split(q, " ")
	a, err := strconv.Atoi(tokens[0])
	if err != nil {
		return 0, false
	}
	res := a
	tokens = tokens[1:]
	for len(tokens) > 0 {
		// accepted operations
		if !acceptedOps[tokens[0]] {
			return 0, false
		}

		// prep
		if tokens[1] == "by" {
			tokens = append(tokens[:1], tokens[2:]...)
		}
		b, err := strconv.Atoi(tokens[1])
		if err != nil {
			return 0, false
		}

		// do the math
		switch tokens[0] {
		case "plus":
			res += b
		case "minus":
			res -= b
		case "multiplied":
			res *= b
		case "divided":
			res /= b
		}

		tokens = tokens[2:]
	}
	return res, true
}
