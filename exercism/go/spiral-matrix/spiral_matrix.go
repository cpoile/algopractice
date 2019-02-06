package spiralmatrix

func SpiralMatrix(size int) [][]int {
	var matrix = make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
	}

	for r, c, d, i := 0, 0, 0, 1; i <= size*size; i++ {
		matrix[r][c] = i
		nextr, nextc := r+dirs[d][0], c+dirs[d][1]
		if !valid(matrix, nextr, nextc) {
			d = (d + 1) % 4
			nextr, nextc = r+dirs[d][0], c+dirs[d][1]
			if !valid(matrix, nextr, nextc) {
				break
			}
		}
		r, c = nextr, nextc
	}
	return matrix
}

var dirs = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func valid(matrix [][]int, r, c int) bool {
	if 0 <= r && r < len(matrix) &&
		0 <= c && c < len(matrix) && matrix[r][c] == 0 {
		return true
	}
	return false
}
