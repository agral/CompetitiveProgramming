package lc2962

import (
	"slices"
)

// Runtime complexity: O(n).
// Auxiliary space: O(1).
// Both are optimal.
func countSubarrays(nums []int, k int) int64 {
	maxValue := slices.Max(nums)
	maxCount := 0
	ans := int64(0)

	l := 0
	for _, num := range nums {
		if num == maxValue {
			maxCount += 1
		}
		// Shrink the window from the left side until maxValue is included k-1 times.
		// Note: no need to bounds check k, it is guaranteed to be at least 1 (1 <= k <= 10^5).
		for maxCount == k {
			if nums[l] == maxValue {
				maxCount -= 1
			}
			l += 1
		}

		ans += int64(l)
	}

	return ans
}
