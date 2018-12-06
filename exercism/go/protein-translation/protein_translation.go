package protein

import (
	"errors"
)

var STOP = errors.New("Stop")
var ErrInvalidBase = errors.New("invalid base")

func FromCodon(s string) (string, error) {
	switch s {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", STOP
	default:
		return "", ErrInvalidBase
	}
}

func FromRNA(s string) ([]string, error) {
	var proteins []string
	for i := 0; i < len(s); i += 3 {
		protein, err := FromCodon(s[i : i+3])
		if err == STOP {
			break
		}
		if err != nil {
			return proteins, err
		}
		proteins = append(proteins, protein)
	}
	return proteins, nil
}
