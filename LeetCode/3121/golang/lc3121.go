package lc3121

// Runtime complexity: O(n)
// Auxiliary space: O(2*26) == O(1)
// Subjective level: easy
// Solved on: 2026-05-26
func numberOfSpecialChars(word string) int {
	const NUM_CHARS = 'Z' - 'A' + 1
	countLowercase := make([]int, NUM_CHARS)
	countUppercase := make([]int, NUM_CHARS)
	for _, ch := range word {
		if ch >= 'A' && ch <= 'Z' {
			countUppercase[ch-'A'] += 1
		} else { // note: it is guaranteed that the string consists of only upper/lowercase English letters.
			// "and every lowercase occurrence of c appears before the first uppercase occurence of c".
			// so if uppercase c has already been seen, mark it as wrong by resetting the lowercase score
			// to zero.
			if countUppercase[ch-'a'] > 0 {
				countLowercase[ch-'a'] = 0
			} else {
				countLowercase[ch-'a'] += 1
			}
		}
	}
	ans := 0
	for offset := range NUM_CHARS {
		if countLowercase[offset] > 0 && countUppercase[offset] > 0 {
			ans += 1
		}
	}

	return ans
}
