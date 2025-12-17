package lc3573

// Runtime complexity: O(n*k)
// Auxiliary space: O(3*n*(k+1)) == O(n*k)
// Subjective level: medium+
// Solved on: 2025-12-17
func maximumProfit(prices []int, k int) int64 {
	n := len(prices)

	// dp[d][t][status]: holds the maximum profit that can be made on first d days,
	// using up to t transactions. Status is one of:
	//  * 0: no stock is held at the end of day d.
	//  * 1: a standard stock is held at the end of day d.
	//  * 2: a stock is shorted at the end of day d.
	dp := make([][][3]int, n)
	for d := range dp {
		dp[d] = make([][3]int, k+1)
	}

	// note: in Golang, dp[0][0][0] is initially 0. No need to change it. Nice.
	for t := 1; t <= k; t++ {
		dp[0][t][1] = -prices[0]
		dp[0][t][2] = prices[0]
	}

	for d := 1; d < n; d++ {
		for t := 1; t <= k; t++ {
			dp[d][t][0] = max(dp[d-1][t][0], prices[d]+dp[d-1][t][1], -prices[d]+dp[d-1][t][2])
			dp[d][t][1] = max(dp[d-1][t][1], -prices[d]+dp[d-1][t-1][0])
			dp[d][t][2] = max(dp[d-1][t][2], prices[d]+dp[d-1][t-1][0])
		}
	}

	return int64(dp[n-1][k][0])
}
