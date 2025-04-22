package lc2338

const MOD int64 = 1_000_000_007

func idealArrays(n int, maxValue int) int {
	// both n and maxValue are constrained to be at most 10^4.
	// Given that 2^14 == 16384 > 10^4, the longest strictly increasing
	// ideal array is: [1, 2, 4, ..., 8192], or in powers of two:
	// [2^0, 2^1, 2^2, ..., 2^13].
	// So the longest possible strictly increasing ideal array has length of 14.
	maxL := min(n, 14)
	factors := getFactors(maxValue)

	// dp[i][j] holds the count of strictly increasing ideal arrays
	// of length i, ending in j. So:
	// * dp[i][0] := sum(dp[i][j]) for j such that 1 <= j <= maxValue,
	// * dp[i][j] := sum(dp[i-1][f]) for f such that j mod f == 0.
	dp := make([][]int64, maxL+1)
	for i := range maxL + 1 {
		dp[i] = make([]int64, maxValue+1)
	}

	// memo is used for memoization; -1 means cell is not memorized yet.
	memo := make([][]int64, n)
	for i := range n {
		memo[i] = make([]int64, maxL)
		for j := range maxL {
			memo[i][j] = -1
		}
		//memo[i] = slices.Repeat([]int64{-1}, maxL)
	}

	for j := 1; j <= maxValue; j++ {
		dp[1][j] = 1
	}
	for i := 2; i <= maxL; i++ {
		for j := 1; j <= maxValue; j++ {
			for _, f := range factors[j] {
				dp[i][j] += dp[i-1][f]
				dp[i][j] %= MOD
			}
		}
	}

	for i := 1; i <= maxL; i++ {
		for j := 1; j <= maxValue; j++ {
			dp[i][0] += dp[i][j]
			dp[i][0] %= MOD
		}
	}

	// Returns the number of distinct ways to create an **ideal** array
	// of size n from a strictly increasing array of length sil
	var calcNumWays func(n int, sil int) int64
	calcNumWays = func(n int, sil int) int64 {
		if sil == 0 {
			return 1
		}
		if sil == n {
			return 1
		}
		if memo[n][sil] != -1 {
			return memo[n][sil]
		}

		memo[n][sil] = (calcNumWays(n-1, sil) + calcNumWays(n-1, sil-1)) % MOD
		return memo[n][sil]
	}

	ans := int64(0)
	for i := 1; i <= maxL; i++ {
		ans += dp[i][0] * calcNumWays(n-1, i-1)
		ans %= MOD
	}

	return int(ans)
}

func getFactors(maxValue int) [][]int64 {
	ans := make([][]int64, maxValue+1)
	for i := 1; i <= maxValue; i++ {
		for j := 2 * i; j <= maxValue; j += i {
			ans[j] = append(ans[j], int64(i))
		}
	}
	return ans
}
