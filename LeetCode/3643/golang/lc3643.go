package lc3643

// Runtime complexity: O(k^2)
// Auxiliary space: O(1) (reusing the original matrix. Otherwise, would have been O(H*W)).
// Subjective level: medium-, just because some genius at Leetcode flipped `x` and `y` parameters.
// Solved on: 2026-03-21
func reverseSubmatrix(grid [][]int, y int, x int, k int) [][]int {
	// NOTE: they fucked up the challenge description. `x` and `y` refer to y and x, respectively.
	// So... to un-fuck it, I've switched the order of the parameters.
	// The original order was:
	// func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int.

	HALF_K := k / 2
	for dh := range HALF_K {
		for w := x; w < x+k; w++ {
			grid[y+dh][w], grid[y+k-1-dh][w] = grid[y+k-1-dh][w], grid[y+dh][w]
		}
	}
	return grid
}
