package lc0056

import (
	"slices"
)

// Runtime complexity: O(sort)
// Auxiliary space: O(n)
// Subjective level: medium. Or once you know the trick, an easy+.
// Solved on: 2026-04-08
func merge(intervals [][]int) [][]int {
	// 1. Sort the intervals in ascending order
	// (w.r.t. the start time; then when tied w.r.t. the end time):
	slices.SortFunc(intervals, func(lhs, rhs []int) int {
		if lhs[0] == rhs[0] {
			return lhs[1] - rhs[1]
		}
		return lhs[0] - rhs[0]
	})

	// 2. Merge the intervals.
	ans := [][]int{}
	ans = append(ans, []int{intervals[0][0], intervals[0][1]})
	lastAnsIdx := 0 // index of the last existing interval in ans slice.
	for i := 1; i < len(intervals); i++ {
		if ans[lastAnsIdx][1] < intervals[i][0] {
			// case 1: this interval does not overlap with the previously processed one.
			// append the current interval verbatim to the answer slice.
			ans = append(ans, []int{intervals[i][0], intervals[i][1]})
			lastAnsIdx += 1
		} else {
			// case 2: this interval overlaps with the previously processed one.
			// extend the previously processed interval, iff the new one's end time
			// exceeds the previously processed's end time.
			ans[lastAnsIdx][1] = max(ans[lastAnsIdx][1], intervals[i][1])
		}
	}

	return ans
}
