package say

// Say implements the extension to use "and" correctly.
// Also, uncomment the exec.Command to have OSX say the number.
func Sayv1(num int64) (string, bool) {
	var ret string
	if num < 0 {
		return ret, false
	} else if num == 0 {
		ret = "zero"
	} else if num < 100 {
		ret = tensv1(num, false)
	} else if num < 1e3 {
		ret = hundredsv1(num, true)
	} else if num < 1e6 {
		ret = thousandsv1(num, true)
	} else if num < 1e9 {
		ret = millionsv1(num, true)
	} else if num < 1e12 {
		ret = billionsv1(num)
	}
	if ret != "" {
		//_ = exec.Command("say", ret).Run()
		return ret, true
	}
	return ret, false
}

var zeroTo19v1 = map[int64]string{
	1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven",
	8: "eight", 9: "nine", 10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen",
	15: "fifteen", 16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen",
}

var twentyTo99v1 = map[int64]string{
	20: "twenty", 30: "thirty", 40: "forty", 50: "fifty", 60: "sixty", 70: "seventy",
	80: "eighty", 90: "ninety",
}

func tensv1(num int64, last bool) (ret string) {
	if last {
		ret += "and "
	}
	if num < 20 {
		return zeroTo19v1[num]
	}
	rem := num % 10
	ret += twentyTo99v1[num-rem]
	if rem > 0 {
		ret += "-" + zeroTo19v1[rem]
	}
	return
}

func hundredsv1(num int64, last bool) (ret string) {
	if num == 0 {
		return
	}
	hundreds := num / 100
	if hundreds > 0 {
		ret += zeroTo19v1[hundreds] + " hundred"
	}
	if rest := tensv1(num%100, last); rest != "" {
		if len(ret) > 0 {
			ret += " "
		}
		ret += rest
	}
	return
}

func thousandsv1(num int64, last bool) (ret string) {
	if num == 0 {
		return
	}
	thousands := num / 1e3
	ret += hundredsv1(thousands, false) + " thousand"
	if rest := hundredsv1(num%1e3, last); rest != "" {
		ret += " " + rest
	}
	return
}

func millionsv1(num int64, last bool) (ret string) {
	if num == 0 {
		return
	}
	millions := num / 1e6
	ret += hundredsv1(millions, false) + " million"
	if rest := thousandsv1(num%1e6, last); rest != "" {
		ret += " " + rest
	}
	return
}

func billionsv1(num int64) (ret string) {
	if num == 0 {
		return
	}
	billions := num / 1e9
	ret += hundredsv1(billions, false) + " billion"
	if rest := millionsv1(num%1e9, true); rest != "" {
		ret += " " + rest
	}
	return
}
