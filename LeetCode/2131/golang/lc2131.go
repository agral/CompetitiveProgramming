package lc2131

// Runtime complexity: O(n)
// Auxiliary space: O(26 * 26) == O(1)
func longestPalindrome(words []string) int {
	const NUM_LETTERS = 26
	ans := 0
	count := make([][]int, NUM_LETTERS)
	for i := range NUM_LETTERS {
		count[i] = make([]int, NUM_LETTERS)
	}

	for _, word := range words {
		f := int(word[0] - 'a')
		s := int(word[1] - 'a')
		if count[s][f] > 0 {
			ans += 4
			count[s][f] -= 1
		} else {
			count[f][s] += 1
		}
	}

	// All possible pairs of two-letter words (4-letters total) have been taken into account
	// at this point. But there may also be one or more single word comprised of two same letters,
	// e.g. "aa" or "zz". If such a word exists, it can be put in the middle of the constructed
	// palindrome; extending its total length by 2. Only one such 2-letter word can be used,
	// even if more than one are available.
	for i := range NUM_LETTERS {
		if count[i][i] > 0 {
			return ans + 2
		}
	}

	return ans
}
