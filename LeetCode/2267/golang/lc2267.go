package lc2267

// Runtime complexity: More than O(mn), less than O(m^2 * n^2).
// Auxiliary space: Same as runtime complexity.
// Subjective level: hard-.
// Solved on: 2026-09-29
func hasValidPath(grid [][]byte) bool {
	//const BALANCE_LIMIT = 100 // can it be optimized/inferred from H,W?
	H := len(grid)
	W := len(grid[0])
	BALANCE_LIMIT := H + W

	// Microoptimization: since we're allowed to go right or down at any step,
	// the total length of the string formed this way is:
	// (H+W-1). It can't be valid if it's of odd length.
	if (H+W)%2 == 0 {
		return false
	}

	// Microoptimization #2: grid[0][0] has to be '(',
	// grid[H-1][W-1] has to be ')':
	if !(grid[0][0] == '(' && grid[H-1][W-1] == ')') {
		return false
	}

	// balance: equals the total count of '(' minus the total count of ')'
	// on the path leading to the cell. Any path resulting in a negative balance
	// yields an invalid solution, so these can be removed immediately - and only
	// paths where balance is non-negative should be taken into account.

	// For every cell: if when reached the number of '(' and ')' even out,
	// then it's still valid (and the exact path-to-here does *not* matter).
	// So it can be memoized without having to account for the previous path.

	// Note: stores true, false and not-yet-known values. So bool is insufficient;
	// using byte values: `0` for false, `1` for true, `42` for unknown.
	dp := make([][][]byte, H)
	for h := range H {
		dp[h] = make([][]byte, W)
		for w := range W {
			dp[h][w] = make([]byte, BALANCE_LIMIT)
			for b := range BALANCE_LIMIT {
				dp[h][w][b] = 42
			}
		}
	}

	// If there's a path from (h,w) to (H-1,W-1) that forms a parentheses string
	// of `balance` balance, returns true. Otherwise returns false.
	var walk func(h, w, balance int) bool
	walk = func(h, w, balance int) bool {
		if h == H || w == W {
			return false // out-of-bounds!
		}

		// handle the balance thing:
		if grid[h][w] == '(' {
			balance += 1
		} else {
			balance -= 1
		}
		// also, the balance can't be negative at any point, as I've written above.
		if balance < 0 {
			return false
		}

		if h == H-1 && w == W-1 {
			// Reached the bottom-right corner; return true iff balanced.
			return balance == 0
		}

		// If already calculated, return the memoized answer.
		if dp[h][w][balance] != 42 {
			return dp[h][w][balance] == 1
		} else { // And if not calculated yet,
			ans := walk(h+1, w, balance) || walk(h, w+1, balance) // calculate it,
			if ans {                                              // memoize it (remember to convert to byte - need a tri-state bool).
				dp[h][w][balance] = 1
			} else {
				dp[h][w][balance] = 0
			}
			return ans // and return the actual true/false verdict.
		}

	}

	return walk(0, 0, 0)
}
