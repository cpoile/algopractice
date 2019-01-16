package beer

import (
	"errors"
	"strconv"
)

func Song() string {
	ret, _ := Verses(99, 0)
	return ret
}

func Verse(v int) (string, error) {
	if v < 0 || v > 99 {
		return "", errors.New("invalid verse number")
	}
	bot1, bot2 := strconv.Itoa(v), strconv.Itoa(v-1)
	firSen := bot1 + " bottles of beer on the wall, " + bot1 + " bottles of beer."
	secSen := "Take one down and pass it around, " + bot2 + " bottles of beer on the wall.\n"
	if v == 2 {
		secSen = "Take one down and pass it around, 1 bottle of beer on the wall.\n"
	}
	if v == 1 {
		firSen = "1 bottle of beer on the wall, 1 bottle of beer."
		secSen = "Take it down and pass it around, no more bottles of beer on the wall.\n"
	}
	if v == 0 {
		firSen = "No more bottles of beer on the wall, no more bottles of beer."
		secSen = "Go to the store and buy some more, 99 bottles of beer on the wall.\n"
	}
	return firSen + "\n" + secSen, nil
}

func Verses(ub, lb int) (string, error) {
	if lb < 0 || ub > 99 || ub < lb {
		return "", errors.New("invalid bounds")
	}
	return recVerses(ub, lb) + "\n", nil
}

func recVerses(ub, lb int) string {
	ret, _ := Verse(ub)
	if ub == lb {
		return ret
	}
	return ret + "\n" + recVerses(ub-1, lb)
}
