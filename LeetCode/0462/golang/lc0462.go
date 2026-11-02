package lc0462

import "slices"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Runtime complexity: O(sort) == O(nlogn)
// Auxiliary space: O(sort)
// Subjective level: medium-
// Solved on: 2026-11-02
func minMoves2(nums []int) int {
	L := len(nums)
	// 1. Sort the array and calculate the median:
	slices.Sort(nums)
	var median int
	if L%2 == 0 { // even-length arrays; midpoint is between (L-1)//2 and (L//2)th element.
		median = (nums[(L-1)/2] + nums[L/2]) / 2
	} else { // odd-length arrays; midpoint is at exactly L/2 (rounded down).
		median = nums[L/2]
	}

	// 2. Minimum number of moves results from incrementing/decrementing each number
	// to the median value of the array.
	ans := 0
	for _, num := range nums {
		ans += abs(num - median)
	}

	return ans
}
