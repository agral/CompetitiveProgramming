package lc0326

// Runtime complexity: O(1), there's only one modulo operation executed.
// Auxiliary space: O(1)
func isPowerOfThree(n int) bool {
	// Max power of three that fits in an int: 3^19 == 1'162'261'467
	return n > 0 && 1_162_261_467%n == 0
}
