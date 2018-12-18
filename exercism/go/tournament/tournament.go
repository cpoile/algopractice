// Package tournament provides a function to tally the results of games
// and output a formatted team ranking
package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	name           string
	mp, w, d, l, p int
}

// Tally reads game results and writes a formatted team ranking
func Tally(reader io.Reader, writer io.Writer) error {
	// our slice of team structs
	var teams []team
	// a name->idx map
	tIdx := map[string]int{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) <= 0 || line[0] == '#' || line[0] == '\n' {
			continue
		}
		tokens := strings.Split(line, ";")
		if len(tokens) != 3 {
			return errors.New("input not well formatted")
		}

		for i := 0; i < 2; i++ {
			// find the team struct or make a new one
			idx, ok := tIdx[tokens[i]]
			if !ok {
				idx = len(teams)
				tIdx[tokens[i]] = idx
				teams = append(teams, team{name: tokens[i]})
			}
			team := teams[idx]

			// record the game result
			switch tokens[2] {
			case "draw":
				team.d++
			case "win":
				if i == 0 {
					team.w++
				} else {
					team.l++
				}
			case "loss":
				if i == 0 {
					team.l++
				} else {
					team.w++
				}
			default:
				return errors.New("result column not well formatted")
			}
			team.mp++
			team.p = team.w*3 + team.d
			teams[idx] = team
		}
	}

	// format the output
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].p == teams[j].p {
			return teams[i].name < teams[j].name
		}
		return teams[i].p > teams[j].p
	})
	fmt.Fprintf(writer, "%-31s| MP |  W |  D |  L |  P\n", "Team")
	for _, t := range teams {
		fmt.Fprintf(writer, "%-31s| %2d | %2d | %2d | %2d | %2d\n",
			t.name, t.mp, t.w, t.d, t.l, t.p)
	}
	return nil
}
