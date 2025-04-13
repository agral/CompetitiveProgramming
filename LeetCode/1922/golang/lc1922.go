package lc1922

const MOD int = 1_000_000_007

// Runtime complexity: O(log n).
// Auxiliary space complexity: O(1).
func countGoodNumbers(n int64) int {
	multiplier := 1
	if n%2 == 1 {
		multiplier = 5
	}
	return permuteMod(20, n/2) * multiplier % MOD
}

func permuteMod(x int, n int64) int {
	if n == 0 {
		return 1
	}
	if n%2 == 1 {
		return x * permuteMod(x, n-1) % MOD
	}
	return permuteMod((x*x)%MOD, n/2)
}
