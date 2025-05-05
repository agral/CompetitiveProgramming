package lc0790

// Runtime complexity: O(n)
// Auxiliary space: O(n)
func numTilings(n int) int {
	const MOD int = 1_000_000_007
	dp := make([]int, 1001)
	dp[1] = 1 // tiling: |
	dp[2] = 2 // tiling: either || or =
	dp[3] = 5 // five tilings available, shown in lc0790's description.
	for i := 4; i <= n; i++ {
		dp[i] = (2*dp[i-1] + dp[i-3]) % MOD
	}
	return dp[n]
}
