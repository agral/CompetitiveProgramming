package lc0032

import "slices"

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: hard, but the trick with ')' extension is so nice.
// Solved on: 2026-10-03
func longestValidParentheses(s string) int {
	L := len(s)
	extended := ")" + s

	//dp[i] holds the length of the longest VPS in extended[1 .. i]
	dp := make([]int, L+1)

	for i := 1; i <= L; i++ {
		if extended[i] == ')' && extended[i-1-dp[i-1]] == '(' {
			dp[i] = dp[i-1] + dp[i-2-dp[i-1]] + 2
		}
	}

	return slices.Max(dp)
}
