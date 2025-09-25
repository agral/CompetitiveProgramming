package lc0120

import "slices"

// Runtime complexity: O(n^2).
// Auxiliary space: O(n^2).
// Subjective level: Easy.
func minimumTotal(triangle [][]int) int {
	H := len(triangle)
	minPathHere := make([][]int, H)
	for i := range H {
		minPathHere[i] = make([]int, i+1)
	}

	// calculate the minPathHere for the edges:
	minPathHere[0][0] = triangle[0][0]
	for h := 1; h < H; h++ {
		minPathHere[h][0] = minPathHere[h-1][0] + triangle[h][0]
		minPathHere[h][h] = minPathHere[h-1][h-1] + triangle[h][h]
	}

	// calculate the minPathHere for the remaining middle part of the triangle:
	for h := 2; h < H; h++ {
		for k := 1; k < h; k++ {
			minPathHere[h][k] = triangle[h][k] + min(minPathHere[h-1][k-1], minPathHere[h-1][k])
		}
	}
	return slices.Min(minPathHere[H-1])
}
