package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var cases = []struct {
	edges    [][]string
	ask      []string
	expected bool
}{
	{[][]string{{"a", "b"}, {"b", "c"}}, []string{"a", "b"}, true},
	{[][]string{{"a", "b"}, {"b", "c"}}, []string{"b", "c"}, true},
	{[][]string{{"a", "b"}, {"b", "c"}}, []string{"a", "c"}, false},
	{[][]string{{"a", "b"}, {"b", "c"}}, []string{"c", "a"}, false},
}

func TestAddEdge(t *testing.T) {
	for _, c := range cases {
		for _, e := range c.edges {
			addEdge(e[0], e[1])
		}
		actual := hasEdge(c.ask[0], c.ask[1])
		assert.Equal(t, c.expected, actual, "given: %v", c.edges)
	}
}
