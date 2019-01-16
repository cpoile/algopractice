package rectangles

var count int

// Count finds each top left corner, then searches for: top right, bottom left,
// bottom right, top left, and checks if top left is the same as where we started
func Count(bs []string) int {
	count = 0

	for r := 0; r < len(bs); r++ {
		for c := 0; c < len(bs[0]); c++ {
			if bs[r][c] == '+' {
				findTR(bs, r, c, r, c)
			}
		}
	}
	return count
}

func findTR(bs []string, rTL, cTL, r, c int) {
	for rNext, cNext := r, c+1; cNext < len(bs[0]); cNext++ {
		if bs[rNext][cNext] == '-' {
			// continue
		} else if bs[rNext][cNext] == '+' {
			findBR(bs, rTL, cTL, rNext, cNext)
			// keep going to see if there are more rectangles to the right
		} else {
			// not '-' or '+', so finished
			return
		}
	}
}

func findBR(bs []string, rTL, cTL, r, c int) {
	for rNext, cNext := r+1, c; rNext < len(bs); rNext++ {
		if bs[rNext][cNext] == '|' {
			// continue
		} else if bs[rNext][cNext] == '+' {
			findBL(bs, rTL, cTL, rNext, cNext)
			// keep going...
		} else {
			// not '|' or '+', so finished
			return
		}
	}
}

func findBL(bs []string, rTL, cTL, r, c int) {
	for rNext, cNext := r, c-1; cNext >= 0; cNext-- {
		if bs[rNext][cNext] == '-' {
			// continue
		} else if bs[rNext][cNext] == '+' {
			findTL(bs, rTL, cTL, rNext, cNext)
			// keep going...
		} else {
			// not '-' or '+', so finished
			return
		}
	}
}

func findTL(bs []string, rTL, cTL, r, c int) {
	for rNext, cNext := r-1, c; rNext >= 0; rNext-- {
		if bs[rNext][cNext] == '|' {
			// continue
		} else if bs[rNext][cNext] == '+' {
			if rNext == rTL && cNext == cTL {
				count++
				return
			} else {
				// not our starting point, continue up if we can
			}
		} else {
			// not '-' or '+', so finished
			return
		}
	}
}
