package lc2975

import "slices"

// Runtime complexity: O(sort) + O(H^2 + W^2)
// Auxiliary space: O(sort) + O(H^2)
// where H and W is the size of hFences+2 and vFences+2, respectively.
// Subjective level: medium
// Solved on: 2026-01-16
// Note: similar to 2943, different in the way a square is formed.
func maximizeSquareArea(h int, w int, hFences []int, vFences []int) int {
	const MOD64 int64 = 1_000_000_007
	const NO_SUCH_GAP = -1

	// 1. Add external fences (that are not present in "internal" fences, so that the dataset is complete:
	hFences = append(hFences, 1)
	hFences = append(hFences, h)
	vFences = append(vFences, 1)
	vFences = append(vFences, w)

	// 2. Sort both horizontal and vertical fences:
	slices.Sort(hFences)
	slices.Sort(vFences)

	// 3. Store all the possible gaps that can be made by removing some (or none) horizontal fences in a set:
	hGaps := map[int]bool{}
	for i := 1; i < len(hFences); i++ {
		for j := 0; j < i; j++ {
			hGaps[hFences[i]-hFences[j]] = true
		}
	}

	// 4. Now generate all the possible vertical gaps. Instead of storing these in a set,
	// just check whether a horizontal gap of this size exists - if so, this may be the answer.
	maxGap := NO_SUCH_GAP
	for i := 1; i < len(vFences); i++ {
		for j := 0; j < i; j++ {
			vGap := vFences[i] - vFences[j]
			if hGaps[vGap] {
				maxGap = max(maxGap, vGap)
			}
		}
	}

	// 5. Return the area of a rectangle made of the largest of such matching gaps,
	// divided modulo by 1e9+7.
	if maxGap == NO_SUCH_GAP {
		return -1
	}

	return int(int64(maxGap) * int64(maxGap) % MOD64)
}
