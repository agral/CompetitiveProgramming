package lc1458

import "math"

// Runtime complexity: O(L1*L2)
// Auxiliary space: O(L1*L2)
// Subjective level: medium+
// Solved on: 2026-01-08
func maxDotProduct(nums1 []int, nums2 []int) int {
	Y := len(nums1)
	X := len(nums2)

	// dp[y][x] holds max dot product of all possible subsequences from: nums1[0:y), nums2[0:x).
	// Initially set to minimum possible integer value.
	dp := make([][]int, Y+1)
	for y := range Y + 1 {
		dp[y] = make([]int, X+1)
		for x := range X + 1 {
			dp[y][x] = math.MinInt32
		}
	}

	// Calculate the actual max dot product of all possible subsequences; incrementally:
	// row-by-row, column-by-column in each row.
	for y := range Y {
		for x := range X {
			dp[y+1][x+1] = max(dp[y+1][x], dp[y][x+1], nums1[y]*nums2[x]+max(0, dp[y][x]))
		}
	}
	return dp[Y][X]
}
