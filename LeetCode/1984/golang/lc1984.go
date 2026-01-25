package lc1984

import "slices"

// Runtime complexity: O(sort) == O(nlogn)
// Auxiliary space: O(sort)
// Subjective level: easy
// Solved on: 2026-01-25
func minimumDifference(nums []int, k int) int {
	n := len(nums)

	// 1. Sort the nums slice in ascending order.
	slices.Sort(nums)

	// 2. One way to pick scores is to select all students from 0 to k-1 (0-indexed 1 to k).
	// Note: it is guaranteed that k <= len(nums), so no need to do bounds checking.
	ans := nums[k-1] - nums[0]

	// 3. But there might be more optimal way; check all the ranges from [1,k] to [n-k, n-1].
	for e := k; e < n; e++ { // e denotes the end of the range. Implicit b == e-(k-1) == e-k+1.
		ans = min(ans, nums[e]-nums[e-k+1])
	}

	// 4. Now ans contains the optimal (minimal) difference of scores of freely chosen k students.
	return ans
}
