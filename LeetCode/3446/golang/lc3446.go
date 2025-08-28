package lc3446

import (
	"slices"
)

// Runtime complexity: O(n^2), where n is the size (==width, ==height) of the input array.
// Auxiliary space: O(n); the input array is reused as output.
// Subjective level: medium, had fun solving this. Did not expect this to be that much fun. Recommended!
func sortMatrix(grid [][]int) [][]int {
	n := len(grid)

	// Iterate over the bottom-left triangular part of the array.
	// This has to be sorted by diagonals in reverse.
	for h := 0; h < n; h++ {
		diag := make([]int, n-h)
		for w := 0; w < n-h; w++ {
			//fmt.Printf("h=%d, w=%d, grid[h+w][w]=%d\n", h, w, grid[h+w][w])
			diag[w] = grid[h+w][w]
		}
		slices.Sort(diag)
		slices.Reverse(diag)
		for w := 0; w < n-h; w++ {
			grid[h+w][w] = diag[w]
		}
	}

	// Iterate over the top-right triangular part of the array.
	// This has to be sorted by diagonals.
	for h := 1; h < n; h++ {
		diag := make([]int, n-h)
		for w := 0; w < n-h; w++ {
			//fmt.Printf("h=%d, w=%d, grid[w][w+h]=%d\n", h, w, grid[w][w+h])
			diag[w] = grid[w][w+h]
		}
		slices.Sort(diag)
		for w := 0; w < n-h; w++ {
			grid[w][w+h] = diag[w]
		}
	}

	return grid
}
