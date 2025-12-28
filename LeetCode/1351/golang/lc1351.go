package lc1351

// Runtime complexity: O(n+m), where n, m are the dimensions of the grid.
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2025-12-28
func countNegatives(grid [][]int) int {
	Y := len(grid) // grid is guaranteed to be at least 1x1.
	X := len(grid[0])
	ans := 0
	// 1. Find the index of the first negative number in the first row.
	//    (if there are no negatives, make it equal X, the width of the grid):
	//    (note: this could be done via binary search, making the solution
	//    O(logn+m) - but I don't feel like overengineering an easy task).
	//    (in fact, for matrices up to 100x100, binary search may be slower
	//    than accessing sequentially (brute force)!)
	firstNegativeIdx := X
	for x, val := range grid[0] {
		if val < 0 {
			firstNegativeIdx = x
			break
		}
	}
	ans += X - firstNegativeIdx

	// 2. In every subsequent column, it is guaranteed that grid[y][firstNegativeIdx] and the following
	//    entries are also negative. But values to the left of firstNegativeIdx can also be negative.
	//    Check and include these.
	for y := 1; y < Y; y++ {
		for firstNegativeIdx > 0 && grid[y][firstNegativeIdx-1] < 0 {
			firstNegativeIdx -= 1
		}
		ans += X - firstNegativeIdx
	}

	return ans
}
