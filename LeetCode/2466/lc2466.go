package lc2466

func countGoodStrings(low int, high int, zero int, one int) int {
	const mod = 1_000_000_007

	// dp[i] holds the number of "good" strings of length exactly i.
	dp := make([]int, high+1)
	dp[0] = 1

	// Fill the dp array from the bottoms up:
	for i := 1; i <= high; i++ {
		if i >= zero {
			dp[i] += dp[i-zero]
			dp[i] %= mod
		}
		if i >= one {
			dp[i] += dp[i-one]
			dp[i] %= mod
		}
	}

	// Calculate the answer:
	ans := 0
	for i := low; i <= high; i++ {
		ans += dp[i]
		ans %= mod
	}
	return ans
}
