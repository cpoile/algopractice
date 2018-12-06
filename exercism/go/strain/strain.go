// Package strain provides filter functions.
package strain

// Ints is a type to hold ints
type Ints []int

// Lists is a type to hold slices of ints
type Lists [][]int

// Strings is a type to hold slices of strings
type Strings []string

// Keep implements a filter on Ints
func (l Ints) Keep(fn func(int) bool) Ints {
	var ret Ints
	for _, v := range l {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

// Discard implements a filter-not on Ints
func (l Ints) Discard(fn func(int) bool) Ints {
	var revFn = func(v int) bool {
		return !fn(v)
	}
	return l.Keep(revFn)
}

// Keep implements a filter on Lists
func (l Lists) Keep(fn func([]int) bool) Lists {
	var ret Lists
	for _, v := range l {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

// Keep implements a filter on Strings
func (l Strings) Keep(fn func(string) bool) Strings {
	var ret Strings
	for _, v := range l {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return ret
}
