package lc1877

import "slices"

// Runtime complexity: O(sort)
// Auxiliary space: O(sort)
// Subjective level: easy
// Solved on: 2026-01-24
func minPairSum(nums []int) int {
	slices.Sort(nums)
	ans := -42 // since all elements of `nums` are guaranteed to be positive,
	// initial ans can be any negative number - it will be overwritten in 1st pass.

	L := len(nums) // len(nums) is guaranteed to be even.
	for i := range L / 2 {
		// after sorting, minimum sum requires pairing lowest and highest numbers together.
		ans = max(ans, nums[i]+nums[L-1-i])
	}
	return ans
}
