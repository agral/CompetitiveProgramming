package lc3513

import "math"

// Runtime complexity: O(log(n))
// Auxiliary space: O(1)
// Subjective level: medium-, requires some mental work.
// Solved on: 2026-07-23
func uniqueXorTriplets(nums []int) int {
	L := len(nums)
	if L <= 2 {
		return L
	}

	return 1 << (int(math.Log2(float64(L))) + 1)
}
