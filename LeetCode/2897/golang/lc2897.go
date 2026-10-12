package lc2897

// Runtime complexity: O(n) (or rather, O(30*n) == O(n))
// Auxiliary space: O(n) + O(30) == O(n)
// Subjective level: hard. But rewarding!
// Solved on: 2026-10-12
func maxSum(nums []int, k int) int {
	const MOD int = 1_000_000_007
	const MOD64 int64 = 1_000_000_007

	// 1 <= nums[i] <= 10^9; and 2 ^ 30 == 1073741824 > 10^9
	// So each ANDed / ORed resulting number still has at most 30 bits.
	const MAX_SET_BITS int = 30

	L := len(nums)
	minIndices := make([]int, MAX_SET_BITS)
	bestFittingNums := make([]int, L)

	ans := 0
	for _, num := range nums {
		for b := range MAX_SET_BITS {
			if (num>>b)&1 == 1 { // if b'th bit is set, the `operation` can be fruitfully performed.
				bestFittingNums[minIndices[b]] |= (1 << b)
				minIndices[b] += 1
			}
		}
	}

	// Finally, calculate the squares of k best fitting numbers. Might overflow an int32.
	for i := range k {
		bf64 := int64(bestFittingNums[i])

		// Note: WA01, this does not keep the result MOD-10e9+7.
		// Keeping for posterity, spent some 5 minutes on debugging
		//ans += int((bf64 * bf64) % MOD64)

		// this does:
		ans += int((bf64 * bf64) % MOD64)
		ans %= MOD
	}
	return ans
}
