package lc1200

import "slices"

// Runtime complexity: O(sort)
// Auxiliary space: O(n)
// Subjective level: easy+
// Solved on: 2026-01-26
func minimumAbsDifference(arr []int) [][]int {
	// 1. Sort the array.
	slices.Sort(arr)

	// 2. Assume that the difference between 0th and 1st element is the only answer.
	// Store this difference as well.
	ans := make([][]int, 1)
	ans[0] = []int{arr[0], arr[1]}
	minDiff := arr[1] - arr[0]

	// 3. Iterate over the rest of pairs, calculating their difference.
	for i := 2; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff < minDiff {
			// 3a. If this pair's difference is lower than existing minDiff,
			// replace the stored answer with this pair,
			// and set the minDiff to the new global low - this value.
			ans = [][]int{{arr[i-1], arr[i]}}
			minDiff = diff
		} else if diff == minDiff {
			// 3b. If this pair's difference equals existing minDiff,
			// add this pair to the answer list. Since arr is sorted,
			// this keeps the invariant of ans being in ascending order.
			ans = append(ans, []int{arr[i-1], arr[i]})
		}
		// 3c. If this pair's difference is higher than minDiff, ignore it.
	}

	return ans
}
