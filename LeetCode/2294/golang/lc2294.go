package lc2294

import "slices"

// Runtime complexity: O(sort) + O(n)
// Auxiliary space: O(sort)
func partitionArray(nums []int, k int) int {
	slices.Sort(nums)
	ans := 1
	minNum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]-minNum > k {
			minNum = nums[i]
			ans += 1
		}
	}
	return ans
}
