package rotationalcipher

func RotationalCipher(s string, key int) string {
	var start rune
	var ret []rune
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z':
			start = 'a'
		case r >= 'A' && r <= 'Z':
			start = 'A'
		default:
			ret = append(ret, r)
			continue
		}
		ret = append(ret, (r-start+rune(key))%26+start)
	}
	return string(ret)
}

// or:
func RotationalCipher2(s string, key int) (ret string) {
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			ret += string((r-'a'+rune(key))%26 + 'a')
		} else if r >= 'A' && r <= 'Z' {
			ret += string((r-'A'+rune(key))%26 + 'A')
		} else {
			ret += string(r)
		}
	}
	return
}
