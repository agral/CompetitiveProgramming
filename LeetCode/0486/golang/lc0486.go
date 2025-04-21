package lc0486

// Runtime complexity: O(n^2)
// Auxiliary space: O(n^2)
// where n is the length of nums array; i.e. |nums|
func predictTheWinner(nums []int) bool {
	// dp[i][j] holds the maximum edge Player 1 can get over Player 2
	// in nums[i..j].
	n := len(nums)
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
		dp[i][i] = nums[i]
	}

	for d := 1; d < n; d++ {
		for i := 0; i+d < n; i++ {
			j := i + d
			dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}
	// Note: Player 1 is considered a winner when there's a tie, too.
	return dp[0][n-1] >= 0
}
