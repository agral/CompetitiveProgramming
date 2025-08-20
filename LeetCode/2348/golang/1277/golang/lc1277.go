package lc1277

// Runtime complexity: O(H*W)
// Auxiliary space: O(1)
// Subjective level: medium
func countSquares(matrix [][]int) int {
	W := len(matrix[0])
	for i := range matrix {
		for j := range W {
			// `matrix` is a matrix of solely ones and zeroes.
			// Use it to hold additional data; this reduces aux space requirements to O(1) (from O(HW)).
			if (matrix[i][j] == 1) && i > 0 && j > 0 {
				// this will make a matrix of 2x2 ones become [[1, 1], [1, 2]].
				// similarly, a matrix of 3x3 ones becomes: [[1, 1, 1], [1, 2, 2], [1, 2, 3]]; and so on.
				// With these new values the number of square submatrices
				// equals the total sum of these modified values across the entire `matrix`.
				matrix[i][j] += min(matrix[i-1][j-1], matrix[i-1][j], matrix[i][j-1])
			}
		}
	}
	ans := 0
	for i := range matrix {
		for j := range W {
			ans += matrix[i][j]
		}
	}
	return ans
}
