package lc1356

import (
	"math/bits"
	"slices"
)

// Runtime complexity: O(sort)
// - (this is due to OnesCount64() being a constant time function)
// Auxiliary space: O(sort)
// Subjective level: easy
// Solved on: 2026-02-25
func sortByBits(arr []int) []int {
	slices.SortFunc(arr, func(lhs int, rhs int) int {
		bcl := bits.OnesCount64(uint64(lhs))
		bcr := bits.OnesCount64(uint64(rhs))
		if bcl == bcr {
			return lhs - rhs
		}
		return bcl - bcr
	})
	return arr
}
