package lc2419

import "slices"

// Runtime complexity: O(n)
// Auxiliary space: O(1)
func longestSubarray(nums []int) int {
	// We're looking for a maximum bitwise AND of any subarray of nums.
	// This means that we're looking for a max element of nums first:
	maxValue := slices.Max(nums)
	streak := 0
	ans := 0
	for _, n := range nums {
		if n == maxValue {
			streak += 1
			ans = max(ans, streak)
		} else {
			streak = 0
		}
	}
	return ans
}
