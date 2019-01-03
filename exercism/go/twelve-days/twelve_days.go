// Package twelve provides functions for creating 'The Twelve Days of Christmas' song.
package twelve

import "fmt"

var days = []string{"", "first", "second", "third", "fourth", "fifth", "sixth",
	"seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
var gifts = []string{"", "a Partridge in a Pear Tree",
	"two Turtle Doves", "three French Hens", "four Calling Birds",
	"five Gold Rings", "six Geese-a-Laying", "seven Swans-a-Swimming",
	"eight Maids-a-Milking", "nine Ladies Dancing", "ten Lords-a-Leaping",
	"eleven Pipers Piping", "twelve Drummers Drumming",
}

// Song assembles the entire 'The Twelve Days of Christmas' song.
func Song() string {
	var ret string
	for i := 1; i <= 12; i++ {
		ret += Verse(i) + "\n"
	}
	return ret
}

// Verse provides a single verse (based on the day) for 'The Twelve Days of Christmas' song.
func Verse(day int) string {
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s",
		days[day], buildDay(day))
}

func buildDay(day int) string {
	if day == 1 {
		return gifts[day] + "."
	} else if day == 2 {
		return gifts[day] + ", and " + buildDay(day-1)
	} else {
		return gifts[day] + ", " + buildDay(day-1)
	}
}
