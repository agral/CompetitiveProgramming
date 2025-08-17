package lc0837

// Runtime complexity: O(n)
// Auxiliary space: O(n)
func new21Game(n int, k int, maxPts int) float64 {
	// P == 1 if n >= k + maxPts - 1
	if k == 0 || n >= k+maxPts-1 {
		return 1.0
	}
	ans := 0.0
	dp := make([]float64, n+1) // dp[i] holds the probability to have i points in the end.
	dp[0] = 1.0
	windowSum := dp[0] // sliding window holds P[i-1] + P[i-2] + ... + P[i - maxPts]

	for i := 1; i <= n; i++ {
		dp[i] = windowSum / float64(maxPts)
		if i >= k {
			ans += dp[i]
		} else {
			windowSum += dp[i]
		}
		if i-maxPts >= 0 {
			// remove the last entry so that windowSum keeps track of exactly maxPts entries.
			windowSum -= dp[i-maxPts]
		}
	}
	return ans
}
