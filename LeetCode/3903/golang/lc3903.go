package lc3903

// Runtime complexity: O(n)
// Auxiliary space: O(2n) == O(n)
// Subjective level: easy. But described poorly.
// Solved on: 2026-09-04
func firstStableIndex(nums []int, k int) int {
	L := len(nums)

	// 1. Calculate the arrays of prefix maximums and suffix minimums.
	// Note: these do include the ith element too, as per the problem description.
	prefixMax := make([]int, L)
	prefixMax[0] = nums[0]
	for i := 1; i < L; i++ {
		prefixMax[i] = max(prefixMax[i-1], nums[i])
	}
	suffixMin := make([]int, L)
	suffixMin[L-1] = nums[L-1]
	for i := L - 2; i >= 0; i-- {
		suffixMin[i] = min(suffixMin[i+1], nums[i])
	}

	// 2. Find the first index i for which prefixMax[i] - suffixMin[i] <= k.
	for i := range L {
		if prefixMax[i]-suffixMin[i] <= k {
			return i
		}
	}

	// 3. There was no stable index in nums at all. Return a -1.
	return -1
}
