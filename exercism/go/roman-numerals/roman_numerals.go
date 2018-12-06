// Package romannumerals provides a function to convert an int to a roman numeral.
package romannumerals

import "fmt"

var dToR = [][]string{
	// ones
	{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
	// tens
	{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
	// hundreds
	{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
	// thousands up to 30000
	{"", "M", "MM", "MMM"},
}

// ToRomanNumeral converts an int to a roman numeral. For: 0 < int <= 3000
func ToRomanNumeral(d int) (ret string, err error) {
	if d <= 0 || d > 3000 {
		return ret, fmt.Errorf("not valid input")
	}
	var thous, hunds, tens, ones int
	thous = d / 1000
	hunds = (d % 1000) / 100
	tens = (d % 100) / 10
	ones = d % 10

	ret += dToR[3][thous] + dToR[2][hunds] + dToR[1][tens] + dToR[0][ones]
	return ret, nil
}
