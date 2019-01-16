package foodchain

var objects = []struct{ object, elaboration, exclamation string }{
	{"", "", ""},
	{"fly", "", ""},
	{"spider", " that wriggled and jiggled and tickled inside her", "It wriggled and jiggled and tickled inside her.\n"},
	{"bird", "", "How absurd to swallow a bird!\n"},
	{"cat", "", "Imagine that, to swallow a cat!\n"},
	{"dog", "", "What a hog, to swallow a dog!\n"},
	{"goat", "", "Just opened her throat and swallowed a goat!\n"},
	{"cow", "", "I don't know how she swallowed a cow!\n"},
	{"horse", "", "She's dead, of course!"},
}

func Song() string {
	return Verses(1, 8)
}

func Verse(v int) string {
	ret := "I know an old lady who swallowed a " + objects[v].object +
		".\n" + objects[v].exclamation
	return ret + recLine(v)
}

// recursively construct the verse's lines
func recLine(line int) string {
	if line == 8 {
		return ""
	}
	if line == 1 {
		return "I don't know why she swallowed the fly. Perhaps she'll die."
	}
	return "She swallowed the " + objects[line].object +
		" to catch the " + objects[line-1].object + objects[line-1].elaboration +
		".\n" + recLine(line-1)
}

func Verses(from, toIncl int) string {
	ret := Verse(from)
	for i := from + 1; i <= toIncl; i++ {
		ret += "\n\n" + Verse(i)
	}
	return ret
}
