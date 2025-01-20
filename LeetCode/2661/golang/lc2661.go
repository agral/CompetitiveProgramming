package lc2661

type RowColumn struct {
	row    int
	column int
}

func firstCompleteIndex(arr []int, mat [][]int) int {
	H := len(mat)
	W := len(mat[0])
	// Maps the number in `arr` to its (column, row) position in `mat`.
	// Since this challenge is 1-indexed, all indices are shifted to 0-indexed
	// (e.g. the number "3" will be positioned at rc[2]; "5" at rc[4] and so on).
	rc := make([]RowColumn, len(arr))

	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			rc[mat[h][w]-1] = RowColumn{h, w}
		}
	}

	paintedRows := make([]int, H)
	paintedCols := make([]int, W)
	for ans, num := range arr {
		paintedRows[rc[num-1].row]++
		if paintedRows[rc[num-1].row] == W {
			return ans
		}
		paintedCols[rc[num-1].column]++
		if paintedCols[rc[num-1].column] == H {
			return ans
		}
	}
	return -1 // this line will never be reached for well-formed inputs.
}
