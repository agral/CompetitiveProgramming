package lc1288

import "sort"

// Runtime complexity: O(sort) + O(n) == O(sort)
// Auxiliary space: O(sort) + O(1)
// Subjective level: medium-
// Solved on: 2026-07-06
func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals[:], func(lhs, rhs int) bool {
		if intervals[lhs][0] == intervals[rhs][0] {
			return intervals[lhs][1] >= intervals[rhs][1]
		}
		return intervals[lhs][0] <= intervals[rhs][0]
	})

	// 0 <= l_i < r_i <= 10^5, so initialize the values to negative ones -
	// so that no special treatment of the 1st interval is necessary.
	coveredRangeEnd := -1
	ans := 0
	for i := range intervals {
		if intervals[i][1] > coveredRangeEnd {
			ans += 1
			coveredRangeEnd = intervals[i][1]
		}
		// else, the interval is necessarily all covered (the starting point
		// is guaranteed to be at least as big as already processed, the range is sorted).
		// Do nothing; don't count it as an "unique" interval.
	}
	return ans
}
