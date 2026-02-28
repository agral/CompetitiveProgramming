package lc1680

import "math"

// Runtime complexity: O(1)
// Space complexity: O(1)
func getBitLength(n int) int {
	return 1 + int(math.Log2(float64(n)))
}

// Runtime complexity: O(n); this solution has big constant factor;
// and could be improved (but still O(n) is the minimum)
//
// Auxiliary space: O(1)
// Subjective level: medium
// Solved on: 2026-02-28
func concatenatedBinary(n int) int {
	const MOD int = 1_000_000_007
	ans := 0
	for i := 1; i <= n; i++ {
		ans <<= getBitLength(i)
		ans %= MOD
		ans += i
		ans %= MOD
	}
	return ans
}
