// Package leap provides a simple function for calculating is a year is a leap year.
package leap

// IsLeapYear calculates if a given year is a leap year.
func IsLeapYear(y int) bool {
	if y%400 == 0 || (y%4 == 0 && y%100 != 0) {
		return true
	}
	return false
}
