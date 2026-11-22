package lc2124

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: trivially easy
// Solved on: 2026-11-22
func checkString(s string) bool {
	seenB := false
	for _, ch := range s {
		if ch == 'b' {
			seenB = true
		} else if seenB { // ch has to be 'a', don't bother checking
			return false
		}
	}
	return true
}
