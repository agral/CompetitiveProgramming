package lc1009

// Runtime complexity: O(log2(n))
// Auxiliary space: O(1)
// Subjective level: medium-
// Solved on: 2026-03-11
func bitwiseComplement(n int) int {
	// return n ^ 0xFFFFFFFF won't work, don't want to flip the preceding zeroes.
	// so have to prepare a valid matching_mask (0b0000000..0111111..1); where
	// the number of binary ones exactly matches the "meaningful" (flipped)
	// part of the int.
	matching_mask := 1 // note: can't be zero, must also work for n=0.
	for matching_mask < n {
		matching_mask = (matching_mask << 1) + 1
	}
	return n ^ matching_mask
}
