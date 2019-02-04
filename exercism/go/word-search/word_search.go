package wordsearch

import "errors"

// Solve builds a trie and searches for words in the puzzle.
// This is overkill for such small input sizes,
// but I wanted to trie it (har har)
func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	res := make(map[string][2][2]int)

	// build the trie
	root := NewTrieNode()
	for _, w := range words {
		root = put(root, w)
	}

	// make a matrix out of the puzzle
	M, N := len(puzzle), len(puzzle[0])
	grid := make([][]byte, M)
	for i, s := range puzzle {
		grid[i] = []byte(s)
	}

	// each different direction a word can be found
	// left-right, right-left, downwards, upwards,
	// diagonal-right-up, diag-right-down, diag-left-up, diag-left-down
	direction := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		{-1, 1}, {1, 1}, {-1, -1}, {1, -1}}

	// test each cell to see if a word starts there
	for r := 0; r < M; r++ {
		for c := 0; c < N; c++ {
			for _, d := range direction {
				if word, foundr, foundc := findInDirection(root, grid, r, c, d[0], d[1]); word != "" {
					// the test cases use c, r
					res[word] = [2][2]int{{c, r}, {foundc, foundr}}
				}
			}
		}
	}
	for _, w := range words {
		if _, ok := res[w]; !ok {
			return res, errors.New("word not found")
		}
	}
	return res, nil
}

func findInDirection(t *TrieNode, grid [][]byte, r, c, dirr, dirc int) (string, int, int) {
	if t == nil {
		return "", 0, 0
	}
	if t.val != nil {
		return t.val.(string), r - dirr, c - dirc
	}
	if cur, ok := getAt(grid, r, c); ok {
		return findInDirection(t.next[cur], grid, r+dirr, c+dirc, dirr, dirc)
	} else {
		return "", 0, 0
	}
}

func getAt(grid [][]byte, r, c int) (byte, bool) {
	M, N := len(grid), len(grid[0])
	if 0 <= r && r < M && 0 <= c && c < N {
		return grid[r][c] - 'a', true
	}
	return 0, false
}

type TrieNode struct {
	val  interface{}
	next [26]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{}
}

func put(root *TrieNode, key string) *TrieNode {
	return putInto(root, key, 0)
}

func putInto(t *TrieNode, key string, d int) *TrieNode {
	if t == nil {
		t = NewTrieNode()
	}
	if d == len(key) {
		t.val = key
		return t
	}
	c := key[d] - 'a'
	t.next[c] = putInto(t.next[c], key, d+1)
	return t
}
