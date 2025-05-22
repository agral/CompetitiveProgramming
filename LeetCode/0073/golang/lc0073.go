package lc0073

// Runtime complexity: O(HW)
// Auxiliary space: O(1)
func setZeroes(matrix [][]int) {
	H := len(matrix)
	W := len(matrix[0])

	// The two extra bools are required. This is still O(1) aux space.
	isZeroingFirstCol := false
	isZeroingFirstRow := false
	for y := range H {
		if matrix[y][0] == 0 {
			isZeroingFirstCol = true
		}
	}
	for x := range W {
		if matrix[0][x] == 0 {
			isZeroingFirstRow = true
		}
	}

	// In-place requirement: store the to-be-zeroed indices in the first row / first column, respectively:
	for y := 1; y < H; y++ {
		for x := 1; x < W; x++ {
			if matrix[y][x] == 0 {
				matrix[y][0] = 0
				matrix[0][x] = 0
			}
		}
	}

	// Zero the matrix except the first row/column where there is a zero in the header row/column:
	for y := 1; y < H; y++ {
		for x := 1; x < W; x++ {
			if matrix[y][0] == 0 || matrix[0][x] == 0 {
				matrix[y][x] = 0
			}
		}
	}

	// Finally zero the header row/column, if needed:
	if isZeroingFirstCol {
		for y := range H {
			matrix[y][0] = 0
		}
	}
	if isZeroingFirstRow {
		for x := range W {
			matrix[0][x] = 0
		}
	}
}
