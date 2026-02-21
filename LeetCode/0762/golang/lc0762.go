package lc0762

import "math/bits"

// Runtime complexity: O(right-left+1).
// Auxiliary space: O(1)
// Subjective level: easy+
// Solved on: 2026-02-21
func countPrimeSetBits(left int, right int) int {
	// Since 1 <= left <= right <= 10^6, both left and right are < 1024*1024 == 2^20.
	// Prime numbers between 1 and 20 are as follows:

	// Note: can iterate over a slice in O(8), but for a set it's theoretically O(1).
	//PRIME_COUNTS := []int{2, 3, 5, 7, 11, 13, 17, 19}
	PRIMES := map[int]bool{
		2:  true,
		3:  true,
		5:  true,
		7:  true,
		11: true,
		13: true,
		17: true,
		19: true,
	}
	ans := 0
	for num := uint64(left); num <= uint64(right); num++ {
		numOnes := bits.OnesCount64(num)
		_, isPrime := PRIMES[numOnes]
		if isPrime {
			ans += 1
		}

	}
	return ans
}
