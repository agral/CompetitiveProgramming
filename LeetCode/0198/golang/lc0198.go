package lc0198

// Runtime complexity: O(n)
// Auxiliary space: O(n)
func rob(nums []int) int {
	n := len(nums)
	// Handle a corner case - only one house to rob:
	if n == 1 {
		return nums[0]
	}

	// dp[i] is the maximum amount that can be robbed in total from houses[0] to houses[i].
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[n-1]
}
