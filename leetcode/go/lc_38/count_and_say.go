package lc_38

func countAndSay(n int) string {
	said := "1"
	for i := 1; i < n; i++ {
		said = say(said)
	}
	return said
}

func say(s string) string {
	var cur byte
	count := 0
	var ret []byte
	for i := 0; i < len(s); i++ {
		if cur != byte(0) && cur != s[i] {
			ret = append(ret, '0'+byte(count))
			ret = append(ret, cur)
			cur = byte(0)
			count = 0
		}
		if cur == byte(0) {
			cur = s[i]
		}
		count++
	}
	if count > 0 {
		ret = append(ret, '0'+byte(count))
		ret = append(ret, cur)
	}
	return string(ret)
}
