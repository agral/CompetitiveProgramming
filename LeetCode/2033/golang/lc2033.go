package lc2033

import (
	"slices"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Runtime complexity: O(H*W*log(H*W)) (or generally, O(sort) with length of the sorted array H*W).
// Auxiliary space: O(H*W) + O(sort); should still be O(H*W).
// Subjective level: medium.
// Solved on: 2026-04-28
func minOperations(grid [][]int, x int) int {
	H := len(grid)
	W := len(grid[0])

	// 1. Make a flat (1D) slice containing all the entries from the original grid. Sort it.
	flattened := make([]int, H*W)
	for h := range H {
		for w := range W {
			flattened[h*W+w] = grid[h][w]
		}
	}
	slices.Sort(flattened)

	// 2. Check if the grid is unsolvable. It will be iff there's any entry in flattened
	//    whose value minus the lowest entry is not a multiple of x.
	//    (as then there's no way to make both it and the lowest entry match).
	for f := 1; f < H*W; f++ {
		if (flattened[f]-flattened[0])%x != 0 {
			return -1
		}
	}

	// 3. OK, the grid is solvable. Solve it by making all the elements equal the mid one.
	//    This is guaranteed to require the least amount of applying the `operation`.
	midValue := flattened[H*W/2]
	ans := 0
	for _, value := range flattened {
		ans += abs(value-midValue) / x
	}
	return ans
}
