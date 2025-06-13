package lc2616

import "slices"

// Runtime complexity: O(sort) + O(n*log(k)); where k might be: max(nums) - min(nums); not sure!
// Auxiliary space: O(sort)
func minimizeMax(nums []int, p int) int {
	n := len(nums)
	slices.Sort(nums)
	left := 0
	right := nums[n-1] - nums[0]
	for left < right {
		mid := (left + right) / 2

		pairsCount := 0
		for i := 0; i < n-1; i++ {
			if nums[i+1]-nums[i] <= mid {
				i += 1 // once used, this number can not be used in any other pair.
				pairsCount += 1
			}
		}

		if pairsCount < p {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
