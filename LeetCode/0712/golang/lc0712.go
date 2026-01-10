package lc0712

import "fmt"

// Runtime complexity: O(L1*L2)
// Auxiliary space: O(L1*L2)
// Subjective level: medium+
// Solved on: 2026-01-10
func minimumDeleteSum(s1 string, s2 string) int {
	L1 := len(s1)
	L2 := len(s2)
	fmt.Println(L1, L2)

	// dp[y][x] holds minimum cost to make s1[0 .. y] == s2[0 .. x].
	dp := make([][]int, L1+1)
	for y := 0; y <= L1; y++ {
		dp[y] = make([]int, L2+1)
	}

	// try deleting s1[y-1]:
	for y := 1; y <= L1; y++ {
		dp[y][0] = dp[y-1][0] + int(s1[y-1])
	}

	// try deleting s2[x-1]:
	for x := 1; x <= L2; x++ {
		dp[0][x] = dp[0][x-1] + int(s2[x-1])
	}

	// delete s1[y+1..L1] and s2[x+1..L2]; smartly:
	for y := 1; y <= L1; y++ {
		for x := 1; x <= L2; x++ {
			if s1[y-1] == s2[x-1] {
				// no additional cost, as the letters match:
				dp[y][x] = dp[y-1][x-1]
			} else {
				dp[y][x] = min(dp[y-1][x]+int(s1[y-1]), dp[y][x-1]+int(s2[x-1]))
			}
		}
	}

	return dp[L1][L2]
}
