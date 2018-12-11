package cipher

import (
	"strings"
	"unicode"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type myCipher struct {
	shifts []int
}

func (c *myCipher) Encode(s string) (ret string) {
	var count int
	for _, r := range s {
		r = unicode.ToLower(r)
		if r < 'a' || r > 'z' {
			continue
		}
		ret += shift(r, c.shifts[count%len(c.shifts)], 1)
		count++
	}
	return
}

func (c *myCipher) Decode(s string) (ret string) {
	for i, r := range s {
		ret += shift(r, c.shifts[i%len(c.shifts)], -1)
	}
	return
}

func shift(r rune, distance, dir int) string {
	d := (26 + rune(dir*distance)) % 26
	return string((r-'a'+d)%26 + 'a')
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance < -25 || distance > 25 || distance == 0 {
		return nil
	}
	cipher := myCipher{[]int{distance}}
	return &cipher
}

func NewVigenere(key string) Cipher {
	if key == strings.Repeat("a", len(key)) {
		return nil
	}
	var shifts []int
	for _, r := range key {
		if r < 'a' || r > 'z' {
			return nil
		}
		shifts = append(shifts, int(r-'a'))
	}
	return &myCipher{shifts}
}
