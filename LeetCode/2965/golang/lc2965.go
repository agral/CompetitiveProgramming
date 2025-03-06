package lc2965

// Runtime complexity: O(n^2)
// Auxiliary space complexity: O(n^2),
// where n is the size of the grid, as per problem description.
// then n^2 is the total length of the input grid.

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	count := make([]int, n*n+1)
	count[0] = 1 // Unused since the numbers are in <1, n> range; but Go is zero-indexed.

	for _, row := range grid {
		for _, num := range row {
			count[num] += 1
		}
	}
	missing, repeated := 0, 0
	for i, c := range count {
		if c == 0 {
			missing = i
		} else if c == 2 {
			repeated = i
		}
	}
	return []int{repeated, missing}
}
