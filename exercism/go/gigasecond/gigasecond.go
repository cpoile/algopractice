// Package gigasecond provides a function to calculate when a person reaches their gigasecond birthday.
package gigasecond

import "time"

// AddGigasecond calculates when a person reaches their gigasecond birthday.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
