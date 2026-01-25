package lc1866

// Runtime complexity: O(n*k)
// Auxiliary space: O(n*k)
// Subjective level: hard
// Solved on: 2026-01-25
func rearrangeSticks(n int, k int) int {
	const MOD int = 1_000_000_007

	// 1 <= n <= 1000, 1 <= k <= n.
	dp := make([][]int, n+1) // +1 for 0-indexation.
	for y := range n + 1 {
		dp[y] = make([]int, k+1)
	}

	var calcRearrangeSticks func(n int, k int) int
	calcRearrangeSticks = func(n int, k int) int {
		if k == 0 {
			return 0
		}
		if k == n {
			return 1
		}
		// if already calculated, return the cached value:
		if dp[n][k] != 0 {
			return dp[n][k]
		}

		// otherwise calcuate, cache & return:
		ans := (calcRearrangeSticks(n-1, k-1) + calcRearrangeSticks(n-1, k)*(n-1)) % MOD
		dp[n][k] = ans
		return ans
	}

	return calcRearrangeSticks(n, k)
}
