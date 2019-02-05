package connect

// ResultOf uses a union-find algorithm/data structure to discover who is the winner.
// This is a little overkill.
func ResultOf(b []string) (string, error) {
	// assumption: board is complete
	m, n := len(b), len(b[0])
	N := m * n

	// virt1/2 are virtual nodes connected to the
	// top/bottom (for player 'O') or left/right (player 'X)
	virt1, virt2 := N, N+1

	ufO := newUF(m, n, N+2)
	ufX := newUF(m, n, N+2)

	// for each position, connect it to other positions around it.
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			converted := r*n + c
			player := b[r][c]
			if player == 'O' {
				ufO.inPlay[converted] = true
				ufO.connectAllDirs(r, c)
				// if the O is on the top or bottom rows, connect to the virtual node
				if r == 0 {
					ufO.union(converted, virt1)
				}
				if r == m-1 {
					ufO.union(converted, virt2)
				}
			} else if player == 'X' {
				ufX.inPlay[converted] = true
				ufX.connectAllDirs(r, c)
				// if the X is on the left or right cols, connect to the virtual node
				if c == 0 {
					ufX.union(converted, virt1)
				}
				if c == n-1 {
					ufX.union(converted, virt2)
				}
			}
		}
	}

	// now test if anyone won
	if ufO.connected(virt1, virt2) {
		return "O", nil
	}
	if ufX.connected(virt1, virt2) {
		return "X", nil
	}
	return "", nil
}

// legal moves on a grid are: left, right, down, up, up-right, down-left
var legalMoves = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {-1, 1}, {1, -1}}

// connect this row and col with all directions around it.
func (b *uf) connectAllDirs(row, col int) {
	current := row*b.n + col
	for _, dir := range legalMoves {
		r := row + dir[0]
		c := col + dir[1]
		if 0 <= r && r < b.m && 0 <= c && c < b.n && b.inPlay[r*b.n+c] {
			b.union(current, r*b.n+c)
		}
	}
}

// Below is the union-find data structure and algorithm
type uf struct {
	m, n, N int
	id      []int
	sz      []int
	inPlay  []bool
}

func newUF(m, n, extra int) *uf {
	N := m * n
	var id, sz = make([]int, N+extra), make([]int, N+extra)
	for i := 0; i < N+extra; i++ {
		id[i] = i
		sz[i] = 1
	}
	return &uf{m, n, N, id, sz, make([]bool, N+extra)}
}

func (b *uf) connected(v, w int) bool {
	return b.root(v) == b.root(w)
}

// root with path compression. Overkill, but fun.
func (b *uf) root(i int) int {
	for b.id[i] != i {
		b.id[i] = b.id[b.id[i]]
		i = b.id[i]
	}
	return i
}

// weighted union find. Overkill also.
func (b *uf) union(p, q int) {
	i, j := b.root(p), b.root(q)
	if i == j {
		return
	}
	if b.sz[i] < b.sz[j] {
		b.id[i] = j
		b.sz[j] += b.sz[i]
	} else {
		b.id[j] = i
		b.sz[i] += b.sz[j]
	}
}
