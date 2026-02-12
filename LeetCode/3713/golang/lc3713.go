package lc3713

// Runtime complexity: O(n^2), where n is the number of letters in s
// Auxiliary space: O(26) == O(1)
// Subjective level: medium-
// Solved on: 2026-02-12
func longestBalanced(s string) int {
	L := len(s)
	const NUM_LETTERS int = 'z' - 'a' + 1
	const NOT_SEEN_YET = -1
	// 1. Generate all the substrings of length 2 <= l <= L, starting at 0 <= b < L:
	// Optimization: go from longest-possible to shortest-possible such substrings instead.
	for l := L; l >= 2; l-- {
		for b := 0; b+l <= L; b++ {
			// 2. For each such substring, calculate the number of times each letter appears:
			//fmt.Printf("Considering substring: %s\n", s[b:b+l])
			histogram := [NUM_LETTERS]int{}
			for i := b; i < b+l; i++ {
				histogram[s[i]-'a'] += 1
			}

			// 3. If the histogram has only zeroes and one distinct number
			// (repeated any number of times), the substring is balanced.
			// With optimization from 1., it is guaranteed that this is the longest one.
			isBalanced := true
			unique_freq := NOT_SEEN_YET
			for i := 0; isBalanced && i < NUM_LETTERS; i++ {
				if histogram[i] > 0 {
					if unique_freq == NOT_SEEN_YET {
						unique_freq = histogram[i]
					} else {
						isBalanced = (histogram[i] == unique_freq)
					}
				}
			}

			if isBalanced {
				return l
			}
		}
	}
	return 1
}
