package minesweeper

import "errors"

func (b Board) Count() error {
	h := len(b)
	w := len((b)[0])

	if !b.valid(h, w) {
		return errors.New("invalid board")
	}

	for r := 1; r < h-1; r++ {
		for c := 1; c < w-1; c++ {
			if b[r][c] == ' ' {
				// count mines -- hat-tip to kdobmayer for this elegant trick!
				for k := 0; k < 9; k++ {
					if b[r-1+k/3][c-1+k%3] == '*' {
						if b[r][c] == ' ' {
							b[r][c] = '1'
						} else {
							b[r][c]++
						}
					}
				}
			}
		}
	}
	return nil
}

func (b *Board) valid(h, w int) bool {
	for _, r := range (*b)[1 : h-1] {
		if len(r) != w || r[0] != '|' || r[w-1] != '|' {
			return false
		}
		for _, c := range r[1 : w-1] {
			if c != '*' && c != ' ' {
				return false
			}
		}
	}
	if (*b)[0][0] != '+' || (*b)[0][w-1] != '+' ||
		(*b)[h-1][0] != '+' || (*b)[h-1][w-1] != '+' {
		return false
	}
	for i := 1; i < w-1; i++ {
		if (*b)[0][i] != '-' || (*b)[h-1][i] != '-' {
			return false
		}
	}
	return true
}
