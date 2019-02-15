package topological_sort

import (
	"sort"
)

func topoSort(g map[string][]string) []string {
	var order []string
	marked := make(map[string]bool)

	var dfs func(v string)
	dfs = func(v string) {
		marked[v] = true
		for _, w := range g[v] {
			if !marked[w] {
				dfs(w)
			}
		}
		order = append(order, v)
	}

	var sorted []string
	for k := range g {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)

	for _, v := range sorted {
		if !marked[v] {
			dfs(v)
		}
	}
	return order
}

func breadthFirst(f func(string) []string, worklist []string) {
	seen := make(map[string]bool)
	var next string
	for len(worklist) > 0 {
		next, worklist = worklist[0], worklist[1:]
		if !seen[next] {
			seen[next] = true
			worklist = append(worklist, f(next)...)
		}
	}
}
