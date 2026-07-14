package lc3336

import "slices"

// Runtime complexity: O(1) as the size of input numbers is limited.
// Would be an O(log(min(a, b))) for a general version of this algorithm (ints of unlimited size).
func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Runtime complexity: O(L * max(nums) * max(nums)) -- bottom-up DP 2D.
// Auxiliary space: O(L * max(nums) * max(nums))
// Subjective level: hard
// Solved on: 2026-07-14
func subsequencePairCount(nums []int) int {
	const MOD int = 1_000_000_007
	L := len(nums)
	maxNum := slices.Max(nums)

	// dp[i][x][y] holds the number of disjoint sequences `seq_a`, `seq_b`
	// of nums[0..i), for which `GCD(seq_a) == x`, `GCD(seq_b) == y`.
	dp := make([][][]int, L+1)
	for i := range L + 1 {
		dp[i] = make([][]int, maxNum+1)
		for x := range maxNum + 1 {
			dp[i][x] = make([]int, maxNum+1)
		}
	}
	dp[0][0][0] = 1 // by definition, that would be two (distinct) empty sequences.

	// Calculate the 2D DP array from the bottom up:
	for i := range L {
		for x := 0; x <= maxNum; x++ {
			for y := 0; y <= maxNum; y++ {
				// Option I: pick nums[i] in the first subsequence `seq_a`.
				gcdX := gcd(x, nums[i])
				dp[i+1][gcdX][y] += dp[i][x][y]
				dp[i+1][gcdX][y] %= MOD

				// Option II: pick nums[i] in the second subsequence `seq_b`.
				gcdY := gcd(y, nums[i])
				dp[i+1][x][gcdY] += dp[i][x][y]
				dp[i+1][x][gcdY] %= MOD

				// Option III: do not use nums[i] in either subsequences.
				dp[i+1][x][y] += dp[i][x][y]
				dp[i+1][x][y] %= MOD
			}
		}
	}

	ans := 0
	for n := 1; n <= maxNum; n++ {
		ans += dp[L][n][n]
		ans %= MOD
	}
	return ans
}
