package lc2914

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: Easy / Trivially easy
// Solved on: 2026-11-18
func minChanges(s string) int {
	L := len(s)
	ans := 0

	// 1. Partition s into groups of two. For groups of two that are either "01" or "10"
	// it is necessary to flip one bit - increase the answer by one.
	// When the end of the string is reached, return the answer.
	// The string might be partitioned into other even-length substrings in some cases,
	// but that does not matter - the partitioning-into-two always gives the optimal answer.

	// Note: it is guaranteed that s has even length;
	// so it always can be partitioned into groups of two. I'm not rechecking that guarantee.
	for i := 0; i < L; i += 2 {
		if s[i] != s[i+1] {
			ans += 1
		}
	}
	return ans
}
