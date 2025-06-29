package lc1498

import "slices"

// Runtime complexity: O(sort)
// Auxiliary space: O(sort) + O(n)
func numSubseq(nums []int, target int) int {
	const MOD int = 1_000_000_007
	n := len(nums)

	slices.Sort(nums)

	// powers_of_two[i] == (2 ^ i) % MOD
	powers_of_two := make([]int, n)
	powers_of_two[0] = 1
	for i := 1; i < n; i++ {
		powers_of_two[i] = (powers_of_two[i-1] * 2) % MOD
	}

	ans := 0
	left := 0
	right := n - 1
	for left <= right {
		if nums[left]+nums[right] <= target {
			ans += powers_of_two[right-left]
			ans %= MOD
			left += 1
		} else {
			right -= 1
		}
	}

	return ans
}
