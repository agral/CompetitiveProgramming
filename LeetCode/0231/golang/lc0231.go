package lc0231

// Runtime complexity: O(1)
// Auxiliary space: O(1)
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return (n & (n - 1)) == 0
}
