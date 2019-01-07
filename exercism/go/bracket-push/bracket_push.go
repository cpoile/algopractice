package brackets

func Bracket(s string) (bool, error) {
	var stack []rune
	pairs := map[rune]rune{']': '[', '}': '{', ')': '('}
	for _, r := range s {
		switch r {
		case '{', '[', '(':
			stack = append([]rune{r}, stack...)
		case '}', ']', ')':
			if len(stack) == 0 || stack[0] != pairs[r] {
				return false, nil
			}
			stack = stack[1:]
		}
	}

	return len(stack) == 0, nil
}
