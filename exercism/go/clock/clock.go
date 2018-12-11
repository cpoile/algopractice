// Package clock provides a simple clock without day.
package clock

import (
	"fmt"
)

const (
	minutesPerHour = 60
	minutesPerDay  = minutesPerHour * 24
)

// Clock is a simple clock structure with only hour and minute fields.
type Clock struct {
	hour, minute int
}

func newClock(minutes int) Clock {
	return Clock{minutes / minutesPerHour, minutes % minutesPerHour}
}

// New creates a clock with specified hour and minute.
func New(hour, minute int) Clock {
	mins := (minutesPerDay + (hour*minutesPerHour+minute)%minutesPerDay) % minutesPerDay
	return newClock(mins)
}

// String prints the clock as a 24 hour format.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add adds specified minutes to the clock.
func (c Clock) Add(minutes int) Clock {
	mins := (c.hour*minutesPerHour + c.minute + minutes) % minutesPerDay
	return newClock(mins)
}

// Subtract removes specified minutes from the clock.
func (c Clock) Subtract(minutes int) Clock {
	mins := (minutesPerDay + (c.hour*minutesPerHour+c.minute-minutes)%minutesPerDay) % minutesPerDay
	return newClock(mins)
}
