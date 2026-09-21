package lc3524

// Runtime complexity: O(n^2)
// Auxiliary space: O(n)
// Subjective level: hard-
// Solved on: 2026-09-21
func resultArray(nums []int, k int) []int64 {
	// Warning: this question is very convoluted, or at least
	// is described in a complicated way. But it's just a DP problem in the end.

	// dp[rem] holds the count of subarrays whose product equals rem, modulo k.
	dp := make([]int64, k)
	ans := make([]int64, k) // this one holds the answer, zero-initialized.

	for _, num := range nums {
		rem := num % k
		updatedDp := make([]int64, k)
		updatedDp[rem] = 1

		for i := range k {
			thisRem := int(int64(i) * int64(rem) % int64(k)) // might overflow
			updatedDp[thisRem] += dp[i]
		}

		for i := range k {
			ans[i] += updatedDp[i]
		}

		dp = updatedDp
	}
	return ans
}
