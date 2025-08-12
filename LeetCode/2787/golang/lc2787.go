package lc2787

import "math"

func intPow(a, x int) int {
	return int(math.Pow(float64(a), float64(x)))
}

// Runtime complexity: O(n*logn)
// Auxiliary space: O(n)
func numberOfWays(n int, x int) int {
	const MOD = 1_000_000_007

	// dp[i] holds the number of ways to express i as a sum of powers of x of some unique numbers.
	dp := make([]int, n+1)
	dp[0] = 1 // 0 == 0^x.

	a := 1
	ax := intPow(a, x)
	for ax <= n {
		for i := n; i >= ax; i-- {
			dp[i] += dp[i-ax]
			dp[i] %= MOD
		}
		a += 1
		ax = intPow(a, x)
	}
	return dp[n]
}
