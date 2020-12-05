package array

// Rotate90Counterclockwise rotates a matrix 90 degrees counterclockwise.
// For example, [[1, 2]] becomes [[2], [1]].
func Rotate90Counterclockwise(m [][]int) [][]int {
	rows, cols := len(m), len(m[0])
	r := make([][]int, cols)

	for i := 0; i < cols; i++ {
		r[cols-1-i] = make([]int, rows)
		for j := 0; j < rows; j++ {
			r[cols-1-i][j] = m[j][i]
		}
	}

	return r
}
