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
	name                                    string
	matchesPlayed, won, drawn, lost, points int
}

// Tally reads game results and writes a formatted team ranking
func Tally(reader io.Reader, writer io.Writer) error {
	// our slice of team structs, so we can sort them
	var teams []*team
	// a map to look up team by name
	nameToTeam := map[string]*team{}

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

		team1, team2 := nameToTeam[tokens[0]], nameToTeam[tokens[1]]
		if team1 == nil {
			team1 = &team{name: tokens[0]}
			teams = append(teams, team1)
			nameToTeam[tokens[0]] = team1
		}
		if team2 == nil {
			team2 = &team{name: tokens[1]}
			teams = append(teams, team2)
			nameToTeam[tokens[1]] = team2
		}

		// record the game result
		switch tokens[2] {
		case "draw":
			team1.drawn++
			team2.drawn++
		case "win":
			team1.won++
			team2.lost++
		case "loss":
			team1.lost++
			team2.won++
		default:
			return errors.New("result column not well formatted")
		}
		team1.matchesPlayed++
		team2.matchesPlayed++
		team1.points = team1.won*3 + team1.drawn
		team2.points = team2.won*3 + team2.drawn
	}

	// format the output
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].points == teams[j].points {
			return teams[i].name < teams[j].name
		}
		return teams[i].points > teams[j].points
	})
	fmt.Fprintf(writer, "%-31s| MP |  W |  D |  L |  P\n", "Team")
	for _, t := range teams {
		fmt.Fprintf(writer, "%-31s| %2d | %2d | %2d | %2d | %2d\n",
			t.name, t.matchesPlayed, t.won, t.drawn, t.lost, t.points)
	}
	return nil
}
