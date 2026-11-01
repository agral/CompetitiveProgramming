package lc0453

import "slices"

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy+
// Solved on: 2026-11-01
func minMoves(nums []int) int {
	// In one move, you can increment `n-1` elements of the array by `1`.
	// This is equivalent to decrementing one element of the array by `1`,
	// when the target is to make all elements of the array equal.
	// So, the solution: find the minimum number in the array;
	// the minimum number of `moves` is equal to sum of differences between ith
	// element and the minimum (i.e. we're "decrementing" all the numbers to the minimum).

	minValue := slices.Min(nums)
	ans := 0
	for _, num := range nums {
		ans += (num - minValue)
	}

	return ans
}
