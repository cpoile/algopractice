package accumulate

func Accumulate(input []string, fn func(string) string) []string {
	ret := make([]string, len(input))
	for i, s := range input {
		ret[i] = fn(s)
	}
	return ret
}
