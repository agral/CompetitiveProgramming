package lc2144

import "slices"

// Runtime complexity: O(sort)
// Auxiliary space: O(sort)
// Subjective level: easy
// Solved on: 2026-06-01
func minimumCost(cost []int) int {
	// 1. Sort the cost slice in the descending order:
	slices.Sort(cost)
	slices.Reverse(cost)

	// 2. Calculate the cost. Have to pay for every 1st and 2nd candy, and can take the 3rd for free.
	// This greedy approach results in optimal (minimum) total cost.
	ans := 0
	for i, c := range cost {
		if (i+1)%3 != 0 {
			ans += c
		} // else, take every 3rd candy for free, so do nothing.
	}
	return ans
}
