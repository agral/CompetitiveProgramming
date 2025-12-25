package lc3075

import "slices"

// Runtime complexity: O(sort)
// Auxiliary space: O(sort)
// Subjective level: easy
// Solved on: 2025-12-25
func maximumHappinessSum(happiness []int, k int) int64 {
	ans := int64(0)

	// 1. Sort the children in order of decreasing initial happiness:
	slices.SortFunc(happiness, func(lhs int, rhs int) int {
		return rhs - lhs
	})

	// 2. Keep adding children up to k. Their happiness when being added at nth turn
	//    is: max(0, (initial_happiness-n+1)), for 1-indexed turn number.
	for numTurn := 0; numTurn < k; numTurn++ {
		ans += int64(max(0, happiness[numTurn]-numTurn))
	}
	return ans
}
