package lc3459

// Runtime complexity: O(h*w*4*max(h, w)) === O(max(h, w)*hw)
// Auxiliary space: O(h*w*2*2*4) === O(h*w)
// Subjective level: hard.
func lenOfVDiagonal(grid [][]int) int {
	DIRS := [][]int{{-1, -1}, {-1, 1}, {1, 1}, {1, -1}}
	NUM_DIRS := len(DIRS)
	H := len(grid)
	W := len(grid[0])

	// Memo: H x W x 2 (hasTurned) x 2 (nextNum - 0 or 2) x 4 (NUM_DIRS)
	memo := make([][][][][]int, H)
	for h := range H {
		memo[h] = make([][][][]int, W)
		for w := range W {
			memo[h][w] = make([][][]int, 2)
			for hT := range 2 {
				memo[h][w][hT] = make([][]int, 2)
				for n := range 2 {
					memo[h][w][hT][n] = make([]int, 4)
					for d := range NUM_DIRS {
						memo[h][w][hT][n][d] = -1
					}
				}
			}
		}
	}

	var dfs func(h int, w int, hasTurned int, expectedNum int, dir int) int
	dfs = func(h int, w int, hasTurned int, expectedNum int, dir int) int {
		if h < 0 || h == H || w < 0 || w == W {
			return 0
		}
		if grid[h][w] != expectedNum {
			return 0
		}
		compactNum := max(0, expectedNum-1) // either 0 or 1; when 1 it means that the expectedNum is 2.
		if memo[h][w][hasTurned][compactNum][dir] != -1 {
			return memo[h][w][hasTurned][compactNum][dir]
		}
		nextNum := 0
		if expectedNum == 0 {
			nextNum = 2
		}
		dh, dw := DIRS[dir][0], DIRS[dir][1]
		res := 1 + dfs(h+dh, w+dw, hasTurned, nextNum, dir)

		// try turning if possible:
		if hasTurned == 0 {
			nextDir := (dir + 1) % NUM_DIRS
			nh, nw := DIRS[nextDir][0], DIRS[nextDir][1]
			res = max(res, 1+dfs(h+nh, w+nw, 1, nextNum, nextDir))
		}
		memo[h][w][hasTurned][compactNum][dir] = res
		return res
	}

	ans := 0
	for h := range H {
		for w := range W {
			if grid[h][w] == 1 {
				for d := range DIRS {
					ans = max(ans, 1+dfs(h+DIRS[d][0], w+DIRS[d][1], 0, 2, d))
				}
			}
		}
	}
	return ans
}
