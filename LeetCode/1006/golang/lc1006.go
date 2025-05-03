package lc1006

// Runtime complexity: O(1)
// Auxiliary space: O(1)
func clumsy(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	if n == 3 || n == 4 {
		return n + 3
	}
	nmod := n % 4
	if nmod == 0 {
		return n + 1
	}
	if nmod == 1 || nmod == 2 {
		return n + 2
	}
	return n - 1
}
