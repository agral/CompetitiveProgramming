package lc0827

func paint() int

func largestIsland(grid [][]int) int {
	H := len(grid)
	W := len(grid[0])
	islands := make([]int, 2) // initially islands is a slice containing two zeroes: []int{0, 0}
	areaId := 2

	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			if grid[h][w] == 1 {
				paint() //...?
				islands = append(islands, areaId)
				areaId++
			}
		}
	}

	maxArea := 0

	ans := 0
	return ans
}
