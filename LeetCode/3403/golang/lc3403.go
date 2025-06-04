package lc3403

// Runtime complexity: O(n),
// Auxiliary space: O(n),
// where n == len(word).
func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}

	i, j, k := 0, 1, 0
	size := len(word)
	for j+k < size {
		if word[k+i] == word[k+j] {
			k++
		} else if word[k+i] > word[k+j] {
			// Move forward to s[j+k+1], where there may be a lexicographically larger substring;
			// none such substring exists before s[j+k+1].
			j = j + k + 1
			k = 0
		} else {
			// Can safely skip to s[i+k+1], or s[j]; whichever is greater.
			// No lexicographically larger substring can exist before each of these.
			i = max(j, i+k+1)
			j = i + 1
			k = 0
		}
	}
	base := word[i:]
	maxSize := size + 1 - numFriends
	finalSize := min(len(base), maxSize)
	return base[:finalSize]
}
