package lc0051

// Runtime complexity: O(n! * n)
//
// Auxiliary space: O(len(ans))
// Subjective level: hard+
// Solved on: 2025-12-25 to 2026-01-02
func solveNQueens(n int) [][]string {
	const EMPTY = '.'
	const QUEEN = 'Q'
	ans := make([][]string, 0)
	columns := make([]bool, n)
	diagMain := make([]bool, 2*n-1)
	diagAux := make([]bool, 2*n-1)

	var dfs func(y int, board [][]rune)
	dfs = func(y int, board [][]rune) {
		if y == n {
			// Turn the [][]byte representation of the board to a []string:
			string_representation := make([]string, n)
			for y := range n {
				string_representation[y] = string(board[y])
			}

			ans = append(ans, string_representation)
			return
		}

		for x := range n {
			if !(columns[x] || diagMain[y+x] || diagAux[x-y+n-1]) {
				columns[x] = true
				diagMain[y+x] = true
				diagAux[x-y+n-1] = true
				board[y][x] = QUEEN

				dfs(y+1, board)

				board[y][x] = EMPTY
				diagAux[x-y+n-1] = false
				diagMain[y+x] = false
				columns[x] = false
			}
		}
	}

	// Prepare an empty board:
	board := make([][]rune, n)
	for y := range n {
		board[y] = make([]rune, n)
		for x := range n {
			board[y][x] = EMPTY
		}
	}

	dfs(0, board)

	// if len(ans) == 0, this means that there are no solutions.
	// Return a new empty slice-of-slices-of-string instead of an unfinished [][]string
	// (this is a Golang-specific issue (or rather, not an issue, but a language design decision)):
	/*
		if len(ans) == 0 {
			return [][]string{{}}
		}
	*/

	// WA1: it turns out that I've mis-interpreted the requirements;
	// they want a [][]string{} for no valid answers. So no special treatment above necessary.
	return ans
}
