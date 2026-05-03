package lc0796

import "strings"

// Runtime complexity: O(n)
// Auxiliary space: O(2n) == O(n)
// Subjective level: easy
// Solved on: 2026-05-03
func rotateString(s string, goal string) bool {
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}
