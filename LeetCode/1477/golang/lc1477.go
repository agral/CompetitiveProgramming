package lc1477

import "math"

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: medium+, pretty standard but somewhat tedious.
// Solved on: 2026-09-17
func minSumOfLengths(arr []int, target int) int {
	L := len(arr)
	windowSum := 0
	ans := math.MaxInt32

	// bestSoFar[i] holds the minimum total length of subarrays in [0..i]
	// for which sum equals `target`. Or math.MaxInt32 if there are no such subarrays.
	bestSoFar := make([]int, L)
	for bsf := range bestSoFar {
		bestSoFar[bsf] = math.MaxInt32
	}

	l := 0
	for r := range L {
		windowSum += arr[r]      // expand the window by one element
		for windowSum > target { // immediately shrink the window if necessary
			windowSum -= arr[l]
			l += 1
		}

		if windowSum == target {
			if l > 0 && bestSoFar[l-1] != math.MaxInt32 {
				ans = min(ans, bestSoFar[l-1]+(r-l+1))
			}
			bestSoFar[r] = min(bestSoFar[r], r-l+1)
		}

		if r > 0 { // in every right-trimmed case, the minimum "expands" to the right.
			bestSoFar[r] = min(bestSoFar[r], bestSoFar[r-1])
		}
	}

	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}
