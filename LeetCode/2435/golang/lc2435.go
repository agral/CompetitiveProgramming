package lc2435

// Runtime complexity: O(H*W*K)
// Auxiliary space: O(H*W*K)
// Subjective level: hard
// Solved on: 2025-11-26
func numberOfPaths(grid [][]int, K int) int {
	const MOD int = 1_000_000_007
	H := len(grid)
	W := len(grid[0])
	memo := make([][][]int, H)
	for h := range H {
		memo[h] = make([][]int, W)
		for w := range W {
			memo[h][w] = make([]int, K)
			for k := range K {
				memo[h][w][k] = -1
			}
		}
	}

	// getNumPaths returns the total number of paths from (h, w) to (H-1, W-1),
	// where the path's sum mod k == rem
	var getNumPaths func(h int, w int, rem int) int
	getNumPaths = func(h int, w int, rem int) int {
		//fmt.Printf("getNumPaths(h=%d, w=%d, rem=%d)\n", h, w, rem)
		if h == H || w == W {
			return 0
		}
		if h == H-1 && w == W-1 {
			if (grid[h][w]+rem)%K == 0 {
				return 1
			}
			return 0
		}
		remIdx := rem % K
		//fmt.Printf("rem==%d, remIdx==%d, K=%d\n", rem, remIdx, K)
		if memo[h][w][remIdx] != -1 {
			return memo[h][w][remIdx]
		}

		newRem := (grid[h][w] + rem) % K
		ans := (getNumPaths(h+1, w, newRem) + getNumPaths(h, w+1, newRem)) % MOD
		memo[h][w][remIdx] = ans
		return ans
	}

	return getNumPaths(0, 0, K)
}
