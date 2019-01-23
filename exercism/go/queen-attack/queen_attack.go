package queenattack

import "errors"

func CanQueenAttack(white, black string) (attack bool, err error) {
	if !valid(white) || !valid(black) || white == black {
		return false, errors.New("invalid input")
	}
	wRank := white[0] - 'a' + 1
	wCol := white[1] - '1' + 1
	bRank := black[0] - 'a' + 1
	bCol := black[1] - '1' + 1

	if wRank == bRank || wCol == bCol ||
		wCol-(wRank-bRank) == bCol ||
		wCol+(wRank-bRank) == bCol {
		return true, nil
	}

	return false, nil
}

func valid(pos string) bool {
	if len(pos) != 2 || pos[0] < 'a' || pos[0] > 'h' || pos[1] < '1' || pos[1] > '8' {
		return false
	}
	return true
}
