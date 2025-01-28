package lc2658

func findMaxFish(grid [][]int) int {
	H := len(grid)
	W := len(grid[0])
	// It's yet another DFS problem...
	var dfs func(y, x int) int // Have to declare dfs before defining it; otherwise
	dfs = func(y, x int) int { // it can not be called recursively from its body.
		if y < 0 || y >= H || x < 0 || x >= W || grid[y][x] == 0 {
			return 0
		}
		ans := grid[y][x]
		grid[y][x] = 0 // after visiting, set the cell to 0 so that it won't be processed any more.
		return ans + dfs(y-1, x) + dfs(y+1, x) + dfs(y, x-1) + dfs(y, x+1)
	}
	ans := 0
	for h := range H {
		for w := range W {
			ans = max(ans, dfs(h, w))
		}
	}
	return ans
}
