package reverse

func String(s string) string {
	var ret string
	for _, r := range s {
		ret = string(r) + ret
	}
	return ret
}

// Or:
func String2(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	runes := []rune(s)
	for i := 0; i < (len(runes) / 2); i++ {
		runes[i], runes[len(runes)-1-i] = runes[len(runes)-1-i], runes[i]
	}
	return string(runes)
}
