package lc0191

// Runtime complexity: O(32) == O(1)
// Auxiliary space: O(1)
// Subjective level: easy+
// Solved on: 2026-01-30
func hammingWeight(n int) int {
	ans := 0
	for i := 0; i < 32; i++ {
		if (n>>i)&1 == 1 {
			ans += 1
		}
	}
	return ans
}
