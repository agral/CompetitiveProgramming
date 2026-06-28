package lc1846

import "slices"

// Runtime complexity: O(sort) == O(nlogn)
// Auxiliary space: O(sort); other than that it's O(1) as the original array is reused.
// Subjective level: easy.
// Solved on: 2026-06-28
func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	L := len(arr)
	slices.Sort(arr)
	arr[0] = 1 // this is forced by the first condition in 1846's description.
	for i := 1; i < L; i++ {
		arr[i] = min(arr[i], arr[i-1]+1)
	}
	return arr[len(arr)-1]
}
