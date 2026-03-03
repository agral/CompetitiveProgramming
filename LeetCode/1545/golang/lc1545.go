package lc1545

// Runtime complexity: O(n) (at most n recursive calls, each of which runs in constant amortized time)
// Auxiliary space: O(1)
// Subjective level: medium
// Solved on: 2026-03-03
func findKthBit(n int, k int) byte {
	if n == 1 {
		return '0' // constraint #1, S_1 == 0.
	}

	// S_n = S_{n-1} + '1' + reverse(invert(S_{n-1})).
	// So the answer depends on which part of the above recursive sequence
	// k points to.
	middleIdx := 1 << (n - 1) // holds the position of the middle character in S_n (note: n is 1-indexed.)
	if k == middleIdx {
		return '1'
	} else if k < middleIdx {
		return findKthBit(n-1, k)
	} else {
		correspondingK := 2*middleIdx - k
		correspondingBit := findKthBit(n-1, correspondingK) // but have to flip it from '0' <=> '1':
		if correspondingBit == '0' {
			return '1'
		} else {
			return '0'
		}
	}
}
