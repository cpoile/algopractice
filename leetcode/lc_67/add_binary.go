package lc_67

func addBinary(a string, b string) string {
	var carry uint8
	var ret []byte
	for i := 0; i < len(a) || i < len(b); i++ {
		var aDigit, bDigit uint8
		if i < len(a) {
			aDigit = a[len(a)-1-i] - '0'
		}
		if i < len(b) {
			bDigit = b[len(b)-1-i] - '0'
		}
		digit := (aDigit + bDigit + carry) % 2
		carry = (aDigit + bDigit + carry) / 2
		ret = append([]byte{digit + 48}, ret...)
	}
	if carry > 0 {
		ret = append([]byte{carry + 48}, ret...)
	}
	return string(ret)
}
