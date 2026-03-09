package lc3129

// Runtime complexity: O(zero*one)
// Auxiliary space: O(zero*one)
// Subjective level: hard
// Solved on: 2026-03-09
func numberOfStableArrays(zero int, one int, limit int) int {
	const MOD int64 = 1_000_000_007
	// dp[z][o] holds the total count of stable arrays where:
	// - the last number in the array is zero (dpZero) or one (dpOne).
	// - the count of zeroes is z
	// - the count of ones is o
	dpZero := make([][]int64, zero+1)
	dpOne := make([][]int64, zero+1)
	for z := range zero + 1 {
		dpZero[z] = make([]int64, one+1)
		dpOne[z] = make([]int64, one+1)
	}

	for z := 0; z <= min(zero, limit); z++ {
		dpZero[z][0] = 1
	}

	for o := 0; o <= min(one, limit); o++ {
		dpOne[0][o] = 1
	}

	for z := 1; z <= zero; z++ {
		for o := 1; o <= one; o++ {
			dpZero[z][o] = dpZero[z-1][o] + dpOne[z-1][o]
			if z-limit >= 1 {
				dpZero[z][o] += MOD
				dpZero[z][o] -= dpOne[z-limit-1][o]
			}
			dpZero[z][o] %= MOD

			dpOne[z][o] = dpOne[z][o-1] + dpZero[z][o-1]
			if o-limit >= 1 {
				dpOne[z][o] += MOD
				dpOne[z][o] -= dpZero[z][o-limit-1]
			}
			dpOne[z][o] %= MOD
		}
	}

	ans := (dpZero[zero][one] + dpOne[zero][one]) % MOD
	return int(ans)
}
