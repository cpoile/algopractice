// Package twofer implements the twofer exercise in exercism.
package twofer

import "fmt"

// Return the string "One for <name>[default: you], one for me"
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
