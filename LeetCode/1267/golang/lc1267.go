package lc1267

func countServers(grid [][]int) int {
	H := len(grid)
	W := len(grid[0])
	rowSums := make([]int, H)
	colSums := make([]int, W)
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if grid[h][w] == 1 {
				rowSums[h]++
				colSums[w]++
			}
		}
	}

	ans := 0
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if grid[h][w] == 1 && (rowSums[h] >= 2 || colSums[w] >= 2) {
				ans++
			}
		}
	}
	return ans
}
