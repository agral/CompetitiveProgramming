package lc0048

// Runtime complexity: O(n*n)
// Auxiliary space: O(1) (as required)
// Subjective level: medium. Requires pen and paper to check things out;
// conceptually at most an easy+, but implementing it is quite painful.
// Solved on: 2026-05-04
func rotate(matrix [][]int) {
	N := len(matrix) // matrix is guaranteed to be of size N by N.
	for ring := 0; ring <= N/2; ring++ {
		segment_length := (N - 1) - (2 * ring)
		for s := range segment_length {
			temp := matrix[ring][ring+s]
			matrix[ring][ring+s] = matrix[N-1-ring-s][ring]
			matrix[N-1-ring-s][ring] = matrix[N-1-ring][N-1-ring-s]
			matrix[N-1-ring][N-1-ring-s] = matrix[ring+s][N-1-ring]
			matrix[ring+s][N-1-ring] = temp
		}
	}
}
