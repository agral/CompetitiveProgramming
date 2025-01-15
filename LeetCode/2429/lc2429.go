package lc2429

import "math/bits"

// Runtime complexity: O(2*30) == O(1)
// Aux space requirements: O(1).
func minimizeXor(num1 int, num2 int) int {
	// 1 <= num1, num2 <= 10^9
	// $ echo "obase=2;10^9"|bc -l
	// 111011100110101100101000000000
	// 11 1011 1001 1010 1100 1010 0000 0000 - that's 30 bits
	const MAX_BIT = 30
	availableOnes := bits.OnesCount(uint(num2))

	ans := 0
	// Turn off set bits of num1 as long as there are still bits to use.
	// Note: it is guaranteed that num2 has at least one bit set; it is not zero.
	for i := MAX_BIT - 1; i >= 0; i-- {
		if (num1>>i)&1 == 1 {
			ans |= (1 << i)
			availableOnes--
			if availableOnes == 0 {
				return ans
			}
		}
	}

	// If there are still `availableOnes`, these must be used.
	// Least significant bits should be turned on
	// in order to make the result the smallest possible number.
	for i := 0; i < MAX_BIT; i++ {
		if (num1>>i)&1 == 0 {
			ans |= (1 << i)
			availableOnes--
			if availableOnes == 0 {
				return ans
			}
		}
	}

	// runtime should never reach this point. But not panicking for a LC challenge.
	return ans
}
