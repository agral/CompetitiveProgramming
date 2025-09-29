package lc1039

import "math"

// Runtime complexity: O(n^2)
// Auxiliary space: O(n^2)
// Subjective level: medium/hard
func minScoreTriangulation(values []int) int {
	LEN_VALUES := len(values)
	dp := make([][]int, LEN_VALUES)
	for i := range LEN_VALUES {
		dp[i] = make([]int, LEN_VALUES)
	}

	for y := 2; y < LEN_VALUES; y++ {
		for x := y - 2; x >= 0; x-- {
			dp[y][x] = math.MaxInt
			for z := x + 1; z < y; z++ {
				dp[y][x] = min(dp[y][x], dp[z][x]+dp[y][z]+(values[x]*values[y]*values[z]))
			}
		}
	}
	return dp[LEN_VALUES-1][0]
}
