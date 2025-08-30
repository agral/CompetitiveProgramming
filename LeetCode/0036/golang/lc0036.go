package lc0036

// Runtime complexity: O(9*9) == O(1)
// Auxiliary space: O(3*9*9) == O(1)
// Subjective level: medium
func isValidSudoku(board [][]byte) bool {
	rows := make([][]bool, 9)
	for k := range 9 {
		rows[k] = make([]bool, 9)
	}
	cols := make([][]bool, 9)
	for k := range 9 {
		cols[k] = make([]bool, 9)
	}
	grids := make([][]bool, 9)
	for k := range 9 {
		grids[k] = make([]bool, 9)
	}

	for h := range 9 {
		for w := range 9 {
			if board[h][w] == '.' {
				continue
			}
			// at this point board[h][w] is guaranteed to be a byte between '1' and '9'
			num := int(board[h][w] - '1') // num is in [0, 8] range.

			if rows[h][num] {
				return false
			}
			rows[h][num] = true

			if cols[w][num] {
				return false
			}
			cols[w][num] = true

			gridnum := 3*(h/3) + (w / 3)
			if grids[gridnum][num] {
				return false
			}
			grids[gridnum][num] = true
		}
	}

	return true
}
