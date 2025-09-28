package lc0976

import "slices"

// Runtime complexity: O(sort)
// Auxiliary space: O(sort)
// Subjective level: easy
func largestPerimeter(nums []int) int {
	slices.Sort(nums)
	// it is guaranteed that len(nums) >= 3; so no funny corner cases there.
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i] < nums[i-1]+nums[i-2] {
			return nums[i-2] + nums[i-1] + nums[i]
		}
	}
	return 0
}
