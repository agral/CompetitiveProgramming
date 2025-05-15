package lc2900

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Since 1 <= n <= 100, it's at most O(100) == O(1) anyway.
func getLongestSubsequence(words []string, groups []int) []string {
	ans := []string{}
	lastSeenGroup := 42 // anything that's not 0 nor 1 works.
	for i := range words {
		if lastSeenGroup != groups[i] {
			lastSeenGroup = groups[i]
			ans = append(ans, words[i])
		}
	}
	return ans
}
