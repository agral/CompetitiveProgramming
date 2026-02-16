package lc0190

// Runtime complexity: O(8) == O(1)
// Auxiliary space: O(16) == O(1) (for holding the nibble reversal array)
// Subjective level: easy+. It's an interesting problem regarding the follow-up.
// Solved on: 2026-02-16
func reverseBits(n int) int {
	reversal := []int{
		0,  // 0000 -> 0000
		8,  // 0001 -> 1000
		4,  // 0010 -> 0100
		12, // 0011 -> 1100
		2,  // 0100 -> 0010
		10, // 0101 -> 1010
		6,  // 0110 -> 0110
		14, // 0111 -> 1110
		1,  // 1000 -> 0001
		9,  // 1001 -> 1001
		5,  // 1010 -> 0101
		13, // 1011 -> 1101
		3,  // 1100 -> 0011
		11, // 1101 -> 1011
		7,  // 1110 -> 0111
		15, // 1111 -> 1111
	} // manually made & validated. All the numbers in 0-15 appear each exactly once - so looks valid.

	// `reversal` holds a reversed number of a 4-bit nibble.
	// To reverse bits of 32-bit signed integer one needs to look up the reversal 8 times
	// for each 4-bit nibble making up the 32-bit input number.
	// Given that the input is guaranteed even (ending in 0),
	// this algorithm needs not concern itself with integer overflows / negative numbers etc.
	ans := 0

	for nibble := range 8 {
		mask := (15 << (4 * nibble))
		val := (n & mask) >> (4 * nibble)
		ans <<= 4 // or in other words, ans *= 16
		ans += reversal[val]
		//fmt.Printf("mask: %x, val: %x, reversal: %x\n", mask, val, reversal[val])
	}

	return ans
}

// Follow-up: If this function is called many times, how would you optimize it?
// Ans: It is currently runnning at O(8), reading the reversals from a precomputed slice.
// It can be easily made O(4) - just precompute the reversal slice of each of 256 values
// in a full byte, then use a full byte (0xFF) mask. This should be faster; but of course
// only timing the extensive test suites will tell. The precomputation should be done
// programmaticaly (this one is an excellent candidate :-) ).
