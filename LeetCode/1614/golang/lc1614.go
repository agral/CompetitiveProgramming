package lc1614

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2026-09-28
func maxDepth(s string) int {
	ans := 0
	currentDepth := 0
	for _, ch := range s {
		if ch == '(' {
			currentDepth += 1
		} else if ch == ')' {
			ans = max(currentDepth, ans)
			currentDepth -= 1
		}
	}
	return ans
}
