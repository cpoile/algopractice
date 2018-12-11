package secret

const (
	wink = 1 << iota
	doubleBlink
	closeYourEyes
	jump
	reverse
)

func Handshake(code uint) (ret []string) {
	if code&wink != 0 {
		ret = append(ret, "wink")
	}
	if code&doubleBlink != 0 {
		ret = append(ret, "double blink")
	}
	if code&closeYourEyes != 0 {
		ret = append(ret, "close your eyes")
	}
	if code&jump != 0 {
		ret = append(ret, "jump")
	}
	if code&reverse != 0 {
		for i := len(ret)/2 - 1; i >= 0; i-- {
			opp := len(ret) - 1 - i
			ret[i], ret[opp] = ret[opp], ret[i]
		}
	}
	return
}
