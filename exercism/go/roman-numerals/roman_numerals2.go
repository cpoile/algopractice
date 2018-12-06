package romannumerals

import "fmt"

type dToR2 struct {
	d int
	r string
}

var dict = []dToR2{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral2(d int) (res string, err error) {
	if d <= 0 || d > 3000 {
		return res, fmt.Errorf("illegal argument, out of range")
	}
	for _, item := range dict {
		for d >= item.d {
			res += item.r
			d -= item.d
		}
	}
	return res, nil
}
