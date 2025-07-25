package lc3487

import "slices"

// Runtime complexity: O(n)
// Auxiliary space: O(n) - for the set that holds duplicates' info
func maxSum(nums []int) int {
	mx := slices.Max(nums)
	if mx <= 0 {
		return mx
	}
	set := map[int]bool{}
	ans := 0
	for _, n := range nums {
		if set[n] == false {
			set[n] = true
			if n > 0 {
				ans += n
			}
		}
	}
	return ans
}
