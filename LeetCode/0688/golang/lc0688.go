package lc0688

import (
	"math"
)

// Runtime complexity: O(n^2 * k)
// Auxiliary space: O(n^2)
func knightProbability(n int, k int, row int, column int) float64 {
	// DIRS: possible knight moves as (y, x) offsets from its current position.
	DIRS := [][]int{{1, -2}, {2, -1}, {2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}}

	// dp[y][x]: the probability of the knight to stand on [y][x].
	dp := make([][]float64, n)
	for y := range n {
		dp[y] = make([]float64, n)
	}
	dp[row][column] = 1.0

	for range k {
		nextDp := make([][]float64, n)
		for y := range n {
			nextDp[y] = make([]float64, n)
		}

		for y := range n {
			for x := range n {
				for _, dir := range DIRS {
					nextY := y + dir[0]
					nextX := x + dir[1]
					if nextY >= 0 && nextY < n && nextX >= 0 && nextX < n {
						nextDp[nextY][nextX] += dp[y][x]
					}
				}
			}
		}
		dp = nextDp
	}

	ans := 0.0
	for y := range n {
		for x := range n {
			ans += dp[y][x]
		}
	}
	return ans / math.Pow(8, float64(k))
}
