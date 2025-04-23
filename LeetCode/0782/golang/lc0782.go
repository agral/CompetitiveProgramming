package lc0782

// Runtime complexity: O(SZ^2)
// Auxiliary space: O(1)
func movesToChessboard(board [][]int) int {
	SZ := len(board)
	for i := 1; i < SZ; i++ {
		for j := 1; j < SZ; j++ {
			if board[0][0]^board[i][0]^board[i][j]^board[0][j] == 1 {
				return -1
			}
		}
	}

	rowSum, colSum := 0, 0
	for i := range SZ {
		rowSum += board[0][i]
		colSum += board[i][0]
	}
	HALF_SZ, HALF_INC_SZ := SZ/2, (SZ+1)/2
	if (rowSum != HALF_SZ && rowSum != HALF_INC_SZ) || (colSum != HALF_SZ && colSum != HALF_INC_SZ) {
		return -1
	}

	rowSwaps, colSwaps := 0, 0
	for i := range SZ {
		if board[i][0] == i&1 {
			rowSwaps += 1
		}
		if board[0][i] == i&1 {
			colSwaps += 1
		}
	}

	// Odd-sized chessboards are less lenient wrt. the number of swaps: only one way is valid.
	// Even-sized chessboards are easier. See below for odd-sized oddities' explanation.
	if SZ&1 == 1 {
		// odd-sized board; if there's an odd count of column/row-swaps, then it's better to swap
		// the other colors (there's one less of these), we have just picked the top-left chessboard
		// color suboptimally. But this does not hold for an even count of color swaps.
		if rowSwaps&1 == 1 {
			rowSwaps = SZ - rowSwaps
		}
		if colSwaps&1 == 1 {
			colSwaps = SZ - colSwaps
		}
	} else { // Even-sized board: just pick whichever value is lower; there's no catch.
		rowSwaps = min(rowSwaps, SZ-rowSwaps)
		colSwaps = min(colSwaps, SZ-colSwaps)
	}

	// a single swap fixes two wrong spots at once.
	return (rowSwaps + colSwaps) / 2
}
