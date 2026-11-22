package lc3660

import "math"

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: medium+ / hard- . Satisfying to solve.
// Solved on: 2026-11-22
func maxValue(nums []int) []int {
	L := len(nums)

	// Calculate the maximum prefix array:
	prefixMax := make([]int, L)
	prefixMax[0] = nums[0]
	for i := 1; i < L; i++ {
		prefixMax[i] = max(prefixMax[i-1], nums[i])
	}

	// Calculate the answer, keep track of the minimum suffix seen so far.
	// This can be done in a single backwards sweep.
	ans := make([]int, L) // automatically filled with zeroes.
	suffixMin := math.MaxInt32
	for i := L - 1; i >= 0; i-- {
		if i < L-1 { // if there are elements after this element; i.e. it's not the last one
			if prefixMax[i] > suffixMin {
				ans[i] = ans[i+1] // then use the value from the right (if it's possible to jump forward);
			} else {
				ans[i] = prefixMax[i] // otherwise use the local prefix-max.
			}
		} else { // handle the case of this being the last element in the array; so can't jump forward.
			if prefixMax[i] <= suffixMin {
				ans[i] = prefixMax[i] // all is well, could jump backwards.
			} else {
				// couldn't jump backwards; so the answer here is 0. Don't even need to change it,
				// it's already zero from the beginning.
				//ans[i] = 0
			}
		}
		// After handling ith answer, update the minSuffix - include the current value.
		suffixMin = min(suffixMin, nums[i])
	}

	return ans
}
