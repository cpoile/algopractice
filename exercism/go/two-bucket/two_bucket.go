// twobucket uses a graph to solve the two bucket problem
package twobucket

import (
	"errors"
)

// comb is the unique combination of water in buckets one and two
type comb struct {
	one, two int
}

// graph holds a graph of each comb and the possible combs it can be turned into
type graph struct {
	adj            map[comb][]comb
	marked         map[comb]bool
	oneMax, twoMax int
}

// fill will find all the combs this comb can be turned into
func (g *graph) fill(node comb) {
	// we only need to visit each comb once
	g.marked[node] = true

	// pour from one -> two
	if diff := min(node.one, g.twoMax-node.two); diff > 0 {
		c := comb{node.one - diff, node.two + diff}
		g.adj[node] = append(g.adj[node], c)
		if !g.marked[c] {
			g.fill(c)
		}
	}
	// pour from two -> one
	if diff := min(node.two, g.oneMax-node.one); diff > 0 {
		c := comb{node.one + diff, node.two - diff}
		g.adj[node] = append(g.adj[node], c)
		if !g.marked[c] {
			g.fill(c)
		}
	}
	// empty one
	if node.one > 0 {
		c := comb{0, node.two}
		g.adj[node] = append(g.adj[node], c)
		if !g.marked[c] {
			g.fill(c)
		}
	}
	// empty two
	if node.two > 0 {
		c := comb{node.one, 0}
		g.adj[node] = append(g.adj[node], c)
		if !g.marked[c] {
			g.fill(c)
		}
	}
	// fill one
	if node.one < g.oneMax {
		c := comb{g.oneMax, node.two}
		g.adj[node] = append(g.adj[node], c)
		if !g.marked[c] {
			g.fill(c)
		}
	}
	// fill two
	if node.two < g.twoMax {
		c := comb{node.one, g.twoMax}
		g.adj[node] = append(g.adj[node], c)
		if !g.marked[c] {
			g.fill(c)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// Solve creates a graph of each possible combination (comb) of buckets.
// Then it finds the shortest path to the goal.
func Solve(oneMax, twoMax, goal int, startBucket string) (string, int, int, error) {
	if oneMax <= 0 || twoMax <= 0 || goal <= 0 {
		return "", 0, 0, errors.New("invalid input")
	}

	g := &graph{make(map[comb][]comb), make(map[comb]bool), oneMax, twoMax}

	// initialize the graph. We do this for each call to Solve
	// because the oneMax and twoMax may be unique
	g.fill(comb{0, 0})

	// now do a bfs to find shortest path to the goal
	marked := make(map[comb]bool)
	edgeTo := make(map[comb]comb)
	var start, opposite comb
	if startBucket == "one" {
		start = comb{oneMax, 0}
		opposite = comb{0, twoMax}
	} else if startBucket == "two" {
		start = comb{0, twoMax}
		opposite = comb{oneMax, 0}
	} else {
		return "", 0, 0, errors.New("invalid input")
	}
	// we have to remove the opposite from the possible paths (as per instructions)
	g.adj[opposite] = []comb{}

	// queue for BFS shortest path.
	q := []comb{start}
	marked[start] = true

	for len(q) != 0 {
		v := q[0]
		q = q[1:]

		// check if we're finished
		if v.one == goal {
			return "one", movesTo(edgeTo, v), v.two, nil
		} else if v.two == goal {
			return "two", movesTo(edgeTo, v), v.one, nil
		}

		// not finished
		for _, w := range g.adj[v] {
			if !marked[w] {
				marked[w] = true
				edgeTo[w] = v
				q = append(q, w)
			}
		}
	}
	return "", 0, 0, errors.New("not possible to create goal")
}

func movesTo(edgeTo map[comb]comb, v comb) int {
	var i int
	empty := comb{}
	for ; v != empty; v = edgeTo[v] {
		i++
	}
	return i
}
