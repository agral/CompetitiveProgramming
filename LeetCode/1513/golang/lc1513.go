package lc1513

// Runtime complexity: O(n) (optimal)
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2025-11-16
func numSub(s string) int {
	const MOD int = 1_000_000_007
	ans := 0
	currOnes := 0
	for _, num := range s {
		if num == '1' {
			currOnes += 1
			ans += currOnes
			ans %= MOD
		} else {
			currOnes = 0
		}
	}
	return ans
}
