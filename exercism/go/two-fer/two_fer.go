// Package twofer implements the twofer exercise in exercism.
package twofer

import "fmt"

// ShareWith returns the string "One for <name>, one for me."
// The default for <name> is 'you'.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
