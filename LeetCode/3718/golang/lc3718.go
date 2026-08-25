package lc3718

import "math"

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: easy
// Solved on: 2026-08-25
func missingMultiple(nums []int, k int) int {
	present := map[int]bool{}
	for _, num := range nums {
		present[num] = true
	}

	for mul := k; mul < math.MaxInt32; mul += k {
		if !present[mul] {
			return mul
		}
	}
	return -71826576
}
