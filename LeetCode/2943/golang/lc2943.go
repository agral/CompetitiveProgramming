package lc2943

import "slices"

// Runtime complexity: O(sort)
// Auxiliary space: O(sort)
// Subjective level: medium+
// Solved on: 2026-01-15
func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	// 1. Sort the slices, because of course these can be passed in random order:
	slices.Sort(hBars)
	slices.Sort(vBars)

	// 2. Calculate the maximum continuous gap for both horizontal and vertical bars:
	maxContinuousGap := func(bars []int) int {
		ans := 2
		currGap := 2
		for i := 1; i < len(bars); i++ {
			if bars[i]-bars[i-1] == 1 {
				currGap += 1
			} else {
				currGap = 2
			}
			ans = max(ans, currGap)
		}
		return ans
	}

	// 3. A square needs to be formed, so pick the minimum of the two, and return its square
	// (which represents this square's area):
	side := min(maxContinuousGap(hBars), maxContinuousGap(vBars))
	return side * side
}
