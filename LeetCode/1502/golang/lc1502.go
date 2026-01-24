package lc1502

import "slices"

// Runtime complexity: O(sort)
// Auxiliary space: O(sort)
// Subjective level: easy
// Solved on: 2026-01-24
func canMakeArithmeticProgression(arr []int) bool {
	slices.Sort(arr)
	delta := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != delta {
			return false
		}
	}
	return true
}
