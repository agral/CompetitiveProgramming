package lc1260

// Runtime complexity: O(H*W)
// Auxiliary space: O(H*W) - for holding the answer array. It could in theory be reduced even to O(1).
// Subjective level: easy.
// Solved on: 2026-07-20
func shiftGrid(grid [][]int, k int) [][]int {
	H, W := len(grid), len(grid[0])
	ans := make([][]int, H)
	for h := range H {
		ans[h] = make([]int, W)
	}

	for h := range H {
		for w := range W {
			// Return the grid after shifting k times. With each shift operation,
			// every element takes the place of the next element (modulo the entire array).
			// so k shifts mean that every element takes the place of k-th next element,
			// modulo the entire array.
			// Each element's position in the array can be mapped from 2D (h, w)
			// to 1D id := (H*h)+w. Shifting k positions mean adding k to the id
			// (modulo H*W, i.e. the size of the entire array).
			// Finally, the ID has to be translated back to (h, w) pair of indices.
			id := (W*h + w + k) % (H * W)
			ah, aw := id/W, id%W
			ans[ah][aw] = grid[h][w]
		}
	}
	return ans
}
