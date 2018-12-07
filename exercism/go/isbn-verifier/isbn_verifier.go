package isbn

func IsValidISBN(s string) bool {
	var sum int
	count := 10
	for i, r := range s {
		if r == '-' {
			continue
		}
		if r == 'X' && count == 1 && i == len(s)-1 {
			sum += 10
			count--
			break
		}
		if r < '0' || r > '9' {
			return false
		}
		sum += int(r-'0') * count
		count--
	}
	return count == 0 && sum%11 == 0
}
