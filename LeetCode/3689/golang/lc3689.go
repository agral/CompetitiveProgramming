package lc3689

import "slices"

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: *easy*
// Solved on: 2026-06-09

// It's sufficient to "choose" the entire nums array, k times.
// This is an extremely easy problem, but described as a mid one.
// Maybe there are some further variants of this problem with better constraints
// (e.g. if same subarray can't be selected more than once, the "trick"
// that makes this challenge trivial no longer holds).
func maxTotalValue(nums []int, k int) int64 {
	maxNum := slices.Max(nums)
	minNum := slices.Min(nums)
	ans := int64(maxNum-minNum) * int64(k)
	return ans
}
