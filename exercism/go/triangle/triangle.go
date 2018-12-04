// Package triangle classifies a triangle based on its sides.
package triangle

import "math"

// Kind is the type of triangle.
type Kind int

// These are the kinds of triangles.
const (
	NaT = iota
	Equ
	Iso
	Sca
)

// KindFromSides will return the Kind of triangle given its three sides.
func KindFromSides(a, b, c float64) Kind {
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) ||
		math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) ||
		a+b < c || a+c < b || b+c < a ||
		(a == 0 && b == 0 && c == 0) ||
		(a < 0 || b < 0 || c < 0) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}
