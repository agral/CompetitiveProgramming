package lc2099

import (
	"slices"
)

// Runtime complexity: O(sort)
// Auxiliary space: O(sort)
func maxSubsequence(nums []int, k int) []int {
	n := len(nums)

	// Make another slice; let this one hold both values and original indices:
	valIndex := make([]ValIndex, n)
	for idx, val := range nums {
		valIndex[idx].index = idx
		valIndex[idx].value = val
	}

	// Sort the new slice so that max values are in the front:
	slices.SortFunc(valIndex, func(lhs, rhs ValIndex) int {
		if lhs.value == rhs.value {
			// When values are the same, order the entries by index, ascending:
			return lhs.index - rhs.index
		}
		// Otherwise, put greater values first (descending/nonascending order):
		return -(lhs.value - rhs.value)
	})

	// Take the first k entries from the sorted val-index slice:
	subseq := valIndex[:k]

	// Sort the sub-slice again; so that the indices are ordered ascendingly:
	slices.SortFunc(subseq, func(lhs, rhs ValIndex) int {
		return lhs.index - rhs.index
	})

	// Make the answer slice; this is essentially subseq, without the index info:
	ans := make([]int, k)
	for i := range k {
		ans[i] = subseq[i].value
	}
	return ans
}

type ValIndex struct {
	value int
	index int
}
