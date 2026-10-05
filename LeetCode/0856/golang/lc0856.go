package lc0856

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: medium-
// Solved on: 2026-10-05
func scoreOfParentheses(s string) int {
	ans := 0
	depth := 0
	L := len(s)
	for i := 1; i < L; i++ {
		if s[i-1] == '(' && s[i] == ')' {
			ans += 1 << depth
		}
		if s[i-1] == '(' {
			depth += 1
		} else { // then s[i-1] guaranteed to be ')', don't even check.
			depth -= 1
		}
	}
	return ans
}
