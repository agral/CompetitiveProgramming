package lc3418

import "math"

// Runtime complexity: O(3*H*W) == O(H*W)
// Auxiliary space: O(3*H*W) == O(H*W)
// Subjective level: medium+
// Solved on: 2026-04-02
func maximumAmount(coins [][]int) int {
	// Prepare a dp array, dp[h][w][k] := maximum amount of money the robot can gather
	// by starting at [0,0], travelling to [h,w] and killing exactly k robbers on the way.
	H := len(coins)
	W := len(coins[0])
	INF := math.MinInt32 / 2 // Emulate the infinity as half the minimum (negative) int.
	dp := make([][][]int, H)
	for h := range H {
		dp[h] = make([][]int, W)
		for w := range W {
			dp[h][w] = make([]int, 3)
			for k := range 3 {
				dp[h][w][k] = INF
			}
		}
	}

	// The robot starts in (0, 0) and collects the coins present there:
	dp[0][0][0] = coins[0][0]
	if coins[0][0] < 0 {
		// The robber can be present on a starting square too.
		// Neutralization is always possible here:
		dp[0][0][1] = 0
	}

	// Fill up the rest of the array. Each cell can be accessed either from the left,
	// or from the top neighboring cells. Pick the more optimal (better yield) of the two.
	for h := range H {
		for w := range W {
			for k := range 3 {
				// try getting to (h, w) from the top (h-1, w).
				// Not possible on the top row where h==0.
				if h > 0 {
					// process getting there just collecting money (potentially getting robbed):
					dp[h][w][k] = max(dp[h][w][k], dp[h-1][w][k]+coins[h][w])

					// and process getting there killing the robber (not possible on a no-kill route k=0).
					// (note: not collecting anything, also there might not be a robber there,
					// but it does not matter - a valid answer will be returned anyway in the end.
					// If there's no robber but we behave as we did, we lose money anyway.)
					if k > 0 {
						dp[h][w][k] = max(dp[h][w][k], dp[h-1][w][k-1])
					}
				}

				// and try getting to (h, w) from from the left (h, w-1).
				// Not possible on the leftmost column where w==0.
				if w > 0 {
					// process getting there just collecting money (potentially getting robbed):
					dp[h][w][k] = max(dp[h][w][k], dp[h][w-1][k]+coins[h][w])

					// and process getting there killing the robber:
					if k > 0 {
						dp[h][w][k] = max(dp[h][w][k], dp[h][w-1][k-1])
					}
				}
			}
		}
	}
	return max(dp[H-1][W-1][0], dp[H-1][W-1][1], dp[H-1][W-1][2])
}
