package lc3499

import "strings"

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: medium
// Solved on: 2026-07-21
func maxActiveSectionsAfterTrade(s string) int {
	lengthsOfZeroBlocks := []int{}
	for i, ch := range s {
		if ch == '0' {
			if i > 0 && s[i-1] == '0' { // there might be some way to encode this better... but it works.
				lengthsOfZeroBlocks[len(lengthsOfZeroBlocks)-1] += 1
			} else {
				lengthsOfZeroBlocks = append(lengthsOfZeroBlocks, 1)
			}
		}
	}

	// try converting the neigboring pairs of zero blocks to ones, pick the largest such group.
	maxConvertedZeroes := 0
	L := len(lengthsOfZeroBlocks)
	for i := 1; i < L; i++ {
		maxConvertedZeroes = max(maxConvertedZeroes, lengthsOfZeroBlocks[i-1]+lengthsOfZeroBlocks[i])
	}

	return maxConvertedZeroes + strings.Count(s, "1")
}
