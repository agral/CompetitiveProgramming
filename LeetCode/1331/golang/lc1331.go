package lc1331

import "slices"

// Runtime complexity: O(sort) == O(nlogn)
// Auxiliary space: O(L) for the map entry -> rank
// Subjective level: easy+
// Solved on: 2026-07-12
func arrayRankTransform(arr []int) []int {
	L := len(arr)
	// If an empty array is passed here, just return it.
	// Otherwise we'd get a runtime error when trying to assign `rank[sorted[0]]` key
	// in a map we'd be making soon.
	if L == 0 {
		return arr
	}

	// prepare a copy of the original array, then sort it:
	sorted := make([]int, L)
	copy(sorted, arr)
	slices.Sort(sorted)

	// Now make a map of the sorted entries to their rank:
	rank := make(map[int]int)
	currRank := 1
	rank[sorted[0]] = currRank
	for i := 1; i < L; i++ {
		if sorted[i] > sorted[i-1] {
			currRank += 1
		}
		rank[sorted[i]] = currRank
	}

	// Then modify the original array, replacing the numbers with their rank:
	for i, val := range arr {
		arr[i] = rank[val]
	}

	// job's done, return the original array.
	return arr
}
