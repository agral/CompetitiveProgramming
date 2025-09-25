package lc3195

// Runtime complexity: O(H * W)
// Auxiliary space: O(4) == O(1)
// Subjective level: Easy+
func minimumArea(grid [][]int) int {
	H := len(grid)
	W := len(grid[0])
	left, right, top, bottom := W, 0, H, 0
	// No need to handle the case of a grid of all zeroes;
	// it is guaranteed that `grid` contains at least one "1".
	for h := range H {
		for w := range W {
			if grid[h][w] == 1 {
				top = min(top, h)
				bottom = max(bottom, h)
				left = min(left, w)
				right = max(right, w)
			}
		}
	}

	return (1 + right - left) * (1 + bottom - top)
}
