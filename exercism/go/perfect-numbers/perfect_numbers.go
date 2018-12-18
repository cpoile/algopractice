package perfect

import (
	"errors"
	"math"
)

var ErrOnlyPositive = errors.New("only positive")

type Classification int

const (
	ClassificationAbundant = iota
	ClassificationDeficient
	ClassificationPerfect
)

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}
	if n == 1 {
		return ClassificationDeficient, nil
	}
	var factorSum int64 = 1
	sqrtn := int64(math.Sqrt(float64(n)))
	for i := int64(2); i <= sqrtn; i++ {
		if n%i == 0 {
			factorSum += i
			if i != n/i {
				factorSum += n / i
			}
		}
	}

	switch {
	case factorSum > n:
		return ClassificationAbundant, nil
	case factorSum == n:
		return ClassificationPerfect, nil
	default:
		return ClassificationDeficient, nil
	}
}
