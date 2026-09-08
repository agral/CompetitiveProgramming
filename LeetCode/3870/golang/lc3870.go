package lc3870

// Runtime complexity: O(1)
// Auxiliary space: O(1)
// Subjective level: trivially easy.
// Solved on: 2026-09-08
func countCommas(n int) int {
	// 1 <= n <= 10^5, so 1 <= n <= 100,000.
	// So the answer is just the count of the elements >= 1000,
	// but no less than 0.
	return max(0, n-999)
}
