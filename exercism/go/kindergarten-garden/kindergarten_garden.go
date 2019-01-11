package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

type Garden struct {
	children map[string][]string
}

var plant = map[rune]string{
	'C': "clover", 'G': "grass", 'R': "radishes", 'V': "violets",
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	// check for odd numbers of cups
	if len(diagram)%4 != 2 {
		return nil, errors.New("wrong format")
	}
	d := strings.Split(diagram[1:], "\n")

	c := make([]string, len(children))
	copy(c, children)
	sort.Strings(c)

	child := make(map[string][]string)

	for _, row := range d {
		for i, cup := range row {
			thePlant, ok := plant[cup]
			if !ok {
				return nil, errors.New("invalid cup code")
			}
			// hat tip to DordovskyDmitry: https://exercism.io/tracks/go/exercises/kindergarten-garden/solutions/b6b4c0d59d524f3bb5985f348a24aad6
			kid := c[i/2]
			child[kid] = append(child[kid], thePlant)
		}
	}

	// check for duplicate names
	if len(child) != len(c) {
		return nil, errors.New("invalid input")
	}

	return &Garden{child}, nil
}

func (g *Garden) Plants(child string) (ret []string, ok bool) {
	plants, ok := g.children[child]
	if !ok {
		return nil, false
	}
	return plants, true
}
