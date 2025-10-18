package lc3397

import (
	"math"
	"slices"
)

// Runtime complexity: O(sort) + O(n)
// Auxiliary space: O(sort)
// Subjective level: medium
func maxDistinctElements(nums []int, k int) int {
	ans := 0
	used := math.MinInt32
	slices.Sort(nums)
	for _, num := range nums {
		if num+k > used {
			used = max(used+1, num-k)
			ans += 1
		}
	}
	return ans
}
