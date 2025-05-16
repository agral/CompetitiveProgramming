package lc2901

import "slices"

// Returns the number of positions at which the corresponding characters
// of the two provided strings differ.
// [dirty optimization trick] Assumes that both lhs and rhs are of same length.
func hammingDistance(lhs string, rhs string) int {
	ans := 0
	for i := range lhs {
		if lhs[i] != rhs[i] {
			ans++
		}
	}
	return ans
}

// Runtime complexity: O(sz^2 * k); k must be related to the overall length of words in `words`;
// but can't pinpoint it exactly.
// Auxiliary space: O(sz).
func getWordsInLongestSubsequence(words []string, groups []int) []string {
	sz := len(words) // == len(groups), guaranteed.
	// dp[i] holds the length of the longest subsequence that ends in words[i]:
	dp := make([]int, sz)
	for i := range dp {
		dp[i] = 1
	}

	// best[i] holds the best index of words[i]
	best := make([]int, sz)
	for i := range best {
		best[i] = -1
	}

	ans := []string{}
	for i := 1; i < sz; i++ {
		for j := range i {
			if groups[i] == groups[j] {
				continue
			}
			if len(words[i]) != len(words[j]) {
				continue
			}
			if hammingDistance(words[i], words[j]) != 1 {
				continue
			}
			if dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
				best[i] = j
			}
		}
	}

	// the last index of the subsequence is the one with max score in dp:
	lastIndex := 0
	for i := 1; i < sz; i++ {
		if dp[i] > dp[lastIndex] {
			lastIndex = i
		}
	}

	// Construct the answer in reverse going from last to first entry:
	for lastIndex != -1 {
		ans = append(ans, words[lastIndex])
		lastIndex = best[lastIndex]
	}

	slices.Reverse(ans)
	return ans
}
