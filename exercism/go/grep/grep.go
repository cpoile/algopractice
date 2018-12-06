// Package grep provides a grep search function that implements some of its functionality.
package grep

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// Search implements grep-like functionality
func Search(pattern string, flags, files []string) []string {
	flagSet := make(map[string]bool)
	for _, flag := range flags {
		flagSet[flag] = true
	}

	if flagSet["-x"] {
		pattern = "^" + pattern + "$"
	}
	if flagSet["-i"] {
		pattern = "(?i)" + pattern
	}
	var regpat = regexp.MustCompile(pattern)
	ret := []string{}

	for _, name := range files {
		file, err := os.Open(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not open: %s", name)
			return nil
		}
		defer file.Close()

		input := bufio.NewScanner(file)
		var i int
		for input.Scan() {
			line := input.Text()
			i++
			matched := regpat.MatchString(line)
			if matched && !flagSet["-v"] {
				if flagSet["-l"] {
					ret = append(ret, name)
					break
				}
				if flagSet["-n"] {
					line = fmt.Sprintf("%d:%s", i, line)
				}
				if len(files) > 1 {
					line = name + ":" + line
				}
				ret = append(ret, line)
			}
			if !matched && flagSet["-v"] {
				if len(files) > 1 {
					line = name + ":" + line
				}
				ret = append(ret, line)
			}
		}
	}
	return ret
}
