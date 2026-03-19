package lc3212

// Runtime complexity: O(H*W)
// Auxiliary space: O(H*W)
// Subjective level: medium+
// Solved on: 2026-03-19
func numberOfSubmatrices(grid [][]byte) int {
	H := len(grid)
	W := len(grid[0])
	ans := 0

	// cntX[h][w] holds the total count of 'X's in grid[0..h-1][0..w-1]
	// cntY[h][w] holds the total count of 'Y's.
	cntX := make([][]int, H+1)
	cntY := make([][]int, H+1)
	for h := range H + 1 {
		cntX[h] = make([]int, W+1)
		cntY[h] = make([]int, W+1)
	}

	// make() leaves cntX[0][0] and cntY[0][0] == 0, which is fitting.
	for h := 1; h <= H; h++ {
		for w := 1; w <= W; w++ {
			// cntX and cntY is essentially a prefix sum. But to not count twice we have to subtract too:
			cntX[h][w] = cntX[h-1][w] + cntX[h][w-1] - cntX[h-1][w-1]
			if grid[h-1][w-1] == 'X' {
				cntX[h][w] += 1
			}

			cntY[h][w] = cntY[h-1][w] + cntY[h][w-1] - cntY[h-1][w-1]
			if grid[h-1][w-1] == 'Y' {
				cntY[h][w] += 1
			}

			if cntX[h][w] > 0 && cntX[h][w] == cntY[h][w] {
				ans += 1
			}

			//fmt.Printf("h=%d, w=%d, grid=%c, cntX=%d, cntY=%d, ans=%d\n",
			//	h, w, grid[h-1][w-1], cntX[h][w], cntY[h][w], ans)
		}
	}

	return ans
}
