package lc2966

import "slices"

// Runtime complexity: O(sort) + O(n)
// Auxiliary space: O(sort)
func divideArray(nums []int, k int) [][]int {
	SZ := len(nums)
	slices.Sort(nums)
	ans := make([][]int, 0, SZ/3)
	for i := 0; i < SZ; i += 3 {
		if nums[i+2]-nums[i] > k {
			// Can not divide the numbers into this subarray; so the partition can not be done at all.
			// Return an empty slice.
			return [][]int{}
		}
		ans = append(ans, []int{nums[i], nums[i+1], nums[i+2]})
	}
	return ans
}
