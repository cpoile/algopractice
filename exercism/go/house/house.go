package house

import "strings"

var verses = []string{
	"house that Jack built.",
	"malt\nthat lay in",
	"rat\nthat ate",
	"cat\nthat killed",
	"dog\nthat worried",
	"cow with the crumpled horn\nthat tossed",
	"maiden all forlorn\nthat milked",
	"man all tattered and torn\nthat kissed",
	"priest all shaven and shorn\nthat married",
	"rooster that crowed in the morn\nthat woke",
	"farmer sowing his corn\nthat kept",
	"horse and the hound and the horn\nthat belonged to",
}

func Song() string {
	return assembleSong(1, "")
}

func assembleSong(versenum int, accum string) string {
	if versenum > len(verses) {
		return strings.Trim(accum, "\n\n")
	}
	return assembleSong(versenum+1, accum+Verse(versenum)+"\n\n")
}

func Verse(i int) string {
	return "This is" + assembleVerse(verses[:i])
}

func assembleVerse(verses []string) string {
	if len(verses) == 0 {
		return ""
	}
	return " the " + verses[len(verses)-1] + assembleVerse(verses[:len(verses)-1])
}
