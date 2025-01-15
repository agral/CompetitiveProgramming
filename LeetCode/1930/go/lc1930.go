package lc1930

func countPalindromicSubsequence(s string) int {
	const NUM_LETTERS = 'z' - 'a' + 1
	SZ := len(s)

	firstOccurrence := make([]int, NUM_LETTERS)
	for i := range firstOccurrence {
		firstOccurrence[i] = SZ
	}
	lastOccurrence := make([]int, NUM_LETTERS) // zero-initialized, which is fine.

	for i := range s {
		offset := s[i] - 'a'
		if i < firstOccurrence[offset] {
			firstOccurrence[offset] = i
		}
		lastOccurrence[offset] = i
	}

	ans := 0
	for i := range NUM_LETTERS {
		if firstOccurrence[i] < lastOccurrence[i] {
			// count unique letters between first and last occurrence (not inclusive).
			// that's how many distinct palindromes starting and ending with i can be made.
			seen := map[uint8]bool{}
			for k := firstOccurrence[i] + 1; k < lastOccurrence[i]; k++ {
				seen[s[k]] = true
			}
			ans += len(seen)
		}
	}

	return ans
}
