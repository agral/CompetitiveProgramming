package lc1871

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: medium+
// Solved on: 2026-05-25
func canReach(s string, minJump int, maxJump int) bool {
	L := len(s)

	count := 0
	dp := make([]bool, L)
	dp[0] = true
	for i := minJump; i < L; i++ {
		//fmt.Printf("processing i=%d\n", i)
		if dp[i-minJump] {
			count += 1
			//fmt.Printf("  dp[i-minJump] == dp[%d] == %t, count=%d\n", i-minJump, dp[i-minJump], count)
		}
		if i-maxJump >= 1 && dp[i-maxJump-1] {
			count -= 1
			//fmt.Printf("  dp[i-maxJump-1] == dp[%d] == %t, count=%d\n", i-maxJump-1, dp[i-maxJump-1], count)
		}

		// Note: the bug was here; used `s[i] == 0`, should have used `s[i] == '0'`.
		// Interestingly, Golang usually has strong type coercion - but not in the case
		// of strings, where individual chars can be compared to a literal 0. Today I learned.
		if count != 0 && s[i] == '0' {
			dp[i] = true
		}
		//fmt.Printf("  finalizing %d, count == %d and s[i] == %c, dp[%d] == %t\n", i, count, s[i], i, dp[i])
	}

	return dp[L-1]
}
