package markdown

import (
	"fmt"
	"strings"
)

// Render translates markdown to HTML
func Render(markdown string) string {
	markdown = strings.Replace(markdown, "__", "<strong>", 1)
	markdown = strings.Replace(markdown, "__", "</strong>", 1)
	markdown = strings.Replace(markdown, "_", "<em>", 1)
	markdown = strings.Replace(markdown, "_", "</em>", 1)

	// the finished html
	var html string

	// to keep track of what level we're in
	var header, list int

	// move through each character and convert if it's a markdown tag
	for pos := 0; pos < len(markdown); pos++ {
		char := markdown[pos]
		switch char {
		case '#':
			for char == '#' {
				header++
				pos++
				char = markdown[pos]
			}
			html += fmt.Sprintf("<h%d>", header)
		case '*':
			if list == 0 {
				html += "<ul>"
			}
			html += "<li>"
			list++
			pos++ // move one extra pos
		case '\n':
			if list > 0 {
				html += "</li>"
			}
			if header > 0 {
				html += fmt.Sprintf("</h%d>", header)
				header = 0
			}
		default:
			html += string(char)
		}
	}

	// finished converting, now close open tags and return
	if header > 0 {
		return html + fmt.Sprintf("</h%d>", header)
	}
	if list > 0 {
		return html + "</li></ul>"
	}
	return "<p>" + html + "</p>"
}
