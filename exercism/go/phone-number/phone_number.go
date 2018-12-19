package phonenumber

import (
	"errors"
	"fmt"
)

func Number(input string) (string, error) {
	var clean []byte
	for _, r := range []byte(input) {
		if r >= '0' && r <= '9' {
			clean = append(clean, r)
		}
	}
	if clean[0] == '1' {
		clean = clean[1:]
	}
	if len(clean) != 10 || clean[0] < '2' || clean[3] < '2' {
		return "", errors.New("invalid phone number")
	}
	return string(clean), nil
}

func AreaCode(input string) (string, error) {
	number, err := Number(input)
	if err != nil {
		return "", err
	}
	return number[:3], nil
}

func Format(input string) (string, error) {
	number, err := Number(input)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", number[:3], number[3:6], number[6:10]), nil
}
