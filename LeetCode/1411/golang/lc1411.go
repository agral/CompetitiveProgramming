package lc1411

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: hard.
// Solved on: 2026-01-03
func numOfWays(n int) int {
	// This is not a DP solution, it's called a Matrix Exponentiation. Interesting bit of trivia.
	const MOD int64 = 1_000_000_007

	// Total count of valid two-color combinations, initially 6: 121, 131, 212, 232, 313, 323
	C2 := int64(6)
	// Total count of valid three-color combinations, initially 6: 123, 132, 213, 231, 312, 321
	C3 := int64(6)

	for i := 1; i < n; i++ {
		nextC2 := int64(3)*C2 + int64(2)*C3
		nextC3 := int64(2) * (C2 + C3)
		C2 = nextC2 % MOD
		C3 = nextC3 % MOD
	}

	return int((C2 + C3) % MOD)
}
