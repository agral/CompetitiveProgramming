package lc0718

// Runtime complexity: O(NM)
// Auxiliary space: O(NM)
// Subjective level: medium/hard
// Solved on: 2025-11-18
func findLength(nums1 []int, nums2 []int) int {
	N := len(nums1)
	M := len(nums2)

	// dp[n][m] := the maximum length of a repeated subarray both in nums1[n..N) and nums2[m..M)
	dp := make([][]int, N+1)
	for n := range N + 1 {
		dp[n] = make([]int, M+1)
	}

	ans := 0
	for n := N - 1; n >= 0; n-- {
		for m := M - 1; m >= 0; m-- {
			if nums1[n] == nums2[m] {
				dp[n][m] = dp[n+1][m+1] + 1
				ans = max(ans, dp[n][m])
			}
		}
	}
	return ans
}
