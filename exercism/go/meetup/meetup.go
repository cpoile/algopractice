package meetup

import (
	"time"
)

type WeekSchedule int

const (
	First = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

// Day returns the numeric date of the plain-language meetup date.
func Day(day WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	tmpDay := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	// move the tmpDay to the start of the 7-day period where we'll
	// find the weekday we're looking for
	if day == Teenth {
		tmpDay = tmpDay.AddDate(0, 0, 12)
	} else if day == Last {
		tmpDay = tmpDay.AddDate(0, 1, -7)
	} else {
		for ; day > 0; day-- {
			tmpDay = tmpDay.AddDate(0, 0, 7)
		}
	}

	// now find the weekday we're looking for
	for tmpDay.Weekday() != weekday {
		tmpDay = tmpDay.AddDate(0, 0, 1)
	}
	return tmpDay.Day()
}
