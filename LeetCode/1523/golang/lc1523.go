package lc1523

// Runtime complexity: O(1)
// Auxiliary space: O(1)
// Subjective level: easy+
// Solved on: 2025-12-07
func countOdds(low int, high int) int {
	return (high+1)/2 - (low / 2)
}
