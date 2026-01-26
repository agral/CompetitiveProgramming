package lc0883

// Runtime complexity: O(HW) == O(H^2)
// Auxiliary space: O(1)
// Subjective level: medium+, very tedious to work this one out.
// Solved on: 2026-01-26
func projectionArea(grid [][]int) int {
	H := len(grid)
	// note: W == H, guaranteed by the problem's constraints.
	// I'm using h everywhere to keep memory usage low.
	// W := H
	ans := 0
	for h := range H {
		maxRowValue := 0
		maxColValue := 0
		for w := range H {
			maxRowValue = max(maxRowValue, grid[h][w])

			// Note: since H == W, can flip indices to simultaneously scan column-wise.
			// If H != W, then I would have to scan w-wise, then h-wise.
			maxColValue = max(maxColValue, grid[w][h])

			// count the unit rectangles from XY-plane projection:
			if grid[h][w] > 0 {
				ans += 1
			}
		}
		// count the unit rectangles from XZ- and YZ-plane projections:
		ans += maxRowValue
		ans += maxColValue
	}
	return ans
}
