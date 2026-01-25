package lc0263

// Runtime complexity: O(logn), where log is log5(n), log3(n) and log2(n),
// until options for whole division are exhausted.
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2026-01-25
func isUgly(n int) bool {
	if n == 0 {
		return false
	}
	for n%5 == 0 {
		n /= 5
	}
	for n%3 == 0 {
		n /= 3
	}
	for n%2 == 0 {
		n /= 2
	}
	return n == 1
}
