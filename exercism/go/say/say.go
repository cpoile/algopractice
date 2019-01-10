package say

import (
	"strings"
)

var zeroTo19 = map[int64]string{
	1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven",
	8: "eight", 9: "nine", 10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen",
	15: "fifteen", 16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen",
}

var twentyTo99 = map[int64]string{
	20: "twenty", 30: "thirty", 40: "forty", 50: "fifty", 60: "sixty", 70: "seventy",
	80: "eighty", 90: "ninety",
}

var suffix = map[int64]string{
	1e3: "thousand",
	1e6: "million",
	1e9: "billion",
}

func recSay(num int64, place int64) (ret string) {
	if place == 0 {
		return
	}
	if suf := num / place; suf > 0 {
		ret += sayHundTens(suf) + " " + suffix[place] + " "
	}
	ret += recSay(num%place, place/1e3)

	return strings.Trim(ret, " ")
}

func sayHundTens(num int64) (ret string) {
	if num >= 100 {
		ret += zeroTo19[num/100] + " " + "hundred "
	}
	tens := num % 100
	if tens < 20 {
		ret += zeroTo19[tens]
	} else {
		rem := tens % 10
		ret += twentyTo99[tens-rem]
		if rem > 0 {
			ret += "-" + zeroTo19[rem]
		}
	}
	return ret
}

func Say(num int64) (string, bool) {
	if num < 0 || num > 1e12-1 {
		return "", false
	}
	if num == 0 {
		return "zero", true
	}
	return recSay(num, 1e9), true
}
