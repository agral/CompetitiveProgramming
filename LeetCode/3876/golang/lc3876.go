package lc3876

import "math"

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: medium+
// Solved on: 2026-09-03
func uniformArray(nums1 []int) bool {
	L := len(nums1)
	hasEvenElements := false
	hasOddElements := false
	minEntry := math.MaxInt32
	for i := range L {
		minEntry = min(minEntry, nums1[i])
		if nums1[i]%2 == 0 {
			hasEvenElements = true
		} else {
			hasOddElements = true
		}
	}

	// Corner case: when there are only even or only odd elements, the answer
	// is trivial - just rewrite nums1 to nums2 using operation#1. Return true.
	if !(hasEvenElements && hasOddElements) {
		return true
	}

	// Otherwise, the array has to be converted to be ODD-PARITY.
	// check if the smallest element / elements (there may be more than one)
	// has the required parity - for the request to be solvable, the minimum has to be odd.
	// In other words, if the minimum is even, there's no way to make the array odd-parity.
	return minEntry%2 == 1
}
