package lc0371

// Runtime complexity: O(32) == O(1)
// Auxiliary space: O(1)
// Subjective level: don't bother. One either knows how to do it, or does not.
// Without some experience with half-adders / full-adders
// (e.g. from https://nandgame.com - recommended!), it seems impossible to work this one out.
// Solved on: 2026-01-31
func getSum(a int, b int) int {
	for a != 0 {
		carry := a & b // store the carry bits for bit-by-bit addition. This works on all 32 bits at once.
		b ^= a         // b = a xor b;

		// as long as there are carry bits present, the carry bits will be re-added
		// in the next iteration of the full-adder algorithm.
		a = carry << 1
	}
	// once no carry bits are present and a finally == 0, b holds the sum of both numbers.
	return b
}
