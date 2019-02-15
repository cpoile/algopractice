package topological_sort

import (
	"fmt"
	"testing"
)

func TestTopoSort(t *testing.T) {
	// prereqs maps computer science courses to their prerequisites.
	var prereqs = map[string][]string{
		"algorithms": {"data structures"},
		"calculus":   {"linear algebra"},
		"compilers": {
			"data structures",
			"formal languages",
			"computer organization",
		},
		"data structures":       {"discrete math"},
		"databases":             {"data structures"},
		"discrete math":         {"intro to programming"},
		"formal languages":      {"discrete math"},
		"networks":              {"operating systems"},
		"operating systems":     {"data structures", "computer organization"},
		"programming languages": {"data structures", "computer organization"},
	}
	expected := []string{"intro to programming",
		"discrete math",
		"data structures",
		"algorithms",
		"linear algebra",
		"calculus",
		"formal languages",
		"computer organization",
		"compilers",
		"databases",
		"operating systems",
		"networks",
		"programming languages"}

	order := topoSort(prereqs)

	if fmt.Sprintf("%q", order) != fmt.Sprintf("%q", expected) {
		t.Errorf("Order and expected not equal. \n%10s: %q\n%10s: %q",
			"Order", order, "Expected", expected)
	}
}
