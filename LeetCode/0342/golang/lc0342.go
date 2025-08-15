package lc0342

// Runtime complexity: O(1)
// Auxiliary space: O(1)
func isPowerOfFour(n int) bool {
	// 1st check: non-positive numbers are certainly NOT powers of four.
	if n <= 0 {
		return false
	}

	// 2nd check: powers of four are also powers of two;
	// verify that there's only one bit set in binary representation of n.
	if n&(n-1) != 0 {
		return false
	}

	// 3rd check: verify that the only set bit is at an even index (0-indexing from the right).
	// Consecutive powers of four have the following binary representation:
	// 4^0: 0b00000001
	// 4^1: 0b00000100
	// 4^2: 0b00010000
	// 4^3: 0b01000000
	// and so on. At this point using a bit mask is sufficient to check whether
	// this power of two is also a power of four.

	mask := 0b01010101010101010101010101010101
	return (n & mask) > 0
}
