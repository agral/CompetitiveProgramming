package lc2563

import "slices"

// Runtime complexity: O(sort) == O(nlogn)
// Aux space complexity: O(sort).
func countFairPairs(nums []int, lower int, upper int) int64 {
	countPairsLessThan := func(target_sum int) int64 {
		ans := int64(0)
		left := 0
		right := len(nums) - 1
		for left < right {
			for left < right && nums[left]+nums[right] > target_sum {
				right -= 1
			}
			ans += int64(right - left)
			left += 1
		}
		return ans
	}
	slices.Sort(nums)
	return countPairsLessThan(upper) - countPairsLessThan(lower-1)
}
