package lc1594

// Runtime complexity: O(H*W)
// Auxiliary space: O(2*H*W) == O(H*W)
// Subjective level: medium / medium+
// Solved on: 2026-03-23
func maxProductPath(grid [][]int) int {
	// Note: since all entries are in range -4..4, all the intermediate multiplication results
	// still fit in an int64. But need to divide modulo at the end.
	// To keep with the spirit of this challenge one could keep dividing modulo at every multiplication.
	const MOD64 int64 = 1_000_000_007
	H := len(grid)
	W := len(grid[0])
	// dpLow: holds the lowest known product from (0,0), to (y,x). This can be (and likely will be) negative.
	// dpHigh: holds the highest known product from (0,0) to (y,x). This can be negative as well!
	dpLow := make([][]int64, H)
	dpHigh := make([][]int64, H)
	for h := range H {
		dpLow[h] = make([]int64, W)
		dpHigh[h] = make([]int64, W)
	}

	// The entries on the top and left border have each only one way to be reached.
	// Fill these up:
	dpLow[0][0] = int64(grid[0][0])
	dpHigh[0][0] = int64(grid[0][0])
	for w := 1; w < W; w++ {
		dpLow[0][w] = dpLow[0][w-1] * int64(grid[0][w])
		dpHigh[0][w] = dpLow[0][w]
	}
	for h := 1; h < H; h++ {
		dpLow[h][0] = dpLow[h-1][0] * int64(grid[h][0])
		dpHigh[h][0] = dpLow[h][0]
	}

	// Each of the remaining entries not on the top/left border can be reached
	// either from the top or from the left. Fill these entries up.
	for h := 1; h < H; h++ {
		for w := 1; w < W; w++ {
			if grid[h][w] >= 0 {
				dpLow[h][w] = min(dpLow[h-1][w], dpLow[h][w-1]) * int64(grid[h][w])
				dpHigh[h][w] = max(dpHigh[h-1][w], dpHigh[h][w-1]) * int64(grid[h][w])
			} else {
				dpLow[h][w] = max(dpHigh[h-1][w], dpHigh[h][w-1]) * int64(grid[h][w])
				dpHigh[h][w] = min(dpLow[h-1][w], dpLow[h][w-1]) * int64(grid[h][w])
			}
		}
	}

	highest := max(dpLow[H-1][W-1], dpHigh[H-1][W-1])
	if highest >= 0 {
		return int(highest % MOD64)
	}
	return -1
}
