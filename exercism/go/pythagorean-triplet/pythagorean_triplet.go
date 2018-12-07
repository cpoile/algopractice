package pythagorean

import "math"

type Triplet [3]int

func Range(min, max int) []Triplet {
	return rangeWithMax(min, max, math.MaxInt32)
}

func rangeWithMax(min, max, sum int) []Triplet {
	var ret []Triplet
	for a := min; a <= max; a++ {
		if a > sum {
			break
		}
		for b := a; b <= max; b++ {
			if a+b > sum {
				break
			}
			for c := b; c <= max; c++ {
				if a+b+c > sum || c*c > a*a+b*b {
					break
				}
				if a*a+b*b == c*c {
					ret = append(ret, Triplet{a, b, c})
				}
			}
		}
	}
	return ret
}

func Sum(p int) []Triplet {
	possibles := rangeWithMax(1, 1000, p)
	var ret []Triplet
	for _, t := range possibles {
		if t[0]+t[1]+t[2] == p {
			ret = append(ret, t)
		}
	}
	return ret
}
