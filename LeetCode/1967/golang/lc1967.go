package lc1967

import "strings"

// Runtime complexity: O(n^2) (as Contains is likely O(n), and there's m patterns to be checked.
// Auxiliary space: O(1).
// Subjective level: easy.
// Solved on: 2026-06-29
func numOfStrings(patterns []string, word string) int {
	ans := 0
	for _, pattern := range patterns {
		if strings.Contains(word, pattern) {
			ans += 1
		}
	}
	return ans
}
