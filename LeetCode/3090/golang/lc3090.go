package lc3090

// Runtime complexity: O(n)
// Auxiliary space: O(26) == O(1)
// Subjective level: medium. Requires a sliding window / two pointers' approach, it's a medium.
// Solved on: 2026-08-14
func maximumLengthSubstring(s string) int {
	const NUM_LETTERS = 'z' - 'a' + 1
	L := len(s)
	freq := [NUM_LETTERS]int{}

	// Note: 2 <= s.length <= 100, so the answer has to be at least 2,
	// since any string of length two is valid.
	ans := 2
	freq[s[0]-'a'] += 1
	freq[s[1]-'a'] += 1

	f, l := 0, 1 // indices of first (f) and last (l) chars that still form a valid substring.
	for true {
		// expand until a triplicate appears:
		hasTriplicate := false
		for !hasTriplicate {
			ans = max(ans, l-f+1)
			l += 1
			if l == L {
				return ans
			}
			freq[s[l]-'a'] += 1
			hasTriplicate = (freq[s[l]-'a'] >= 3)
		}

		// contract until there's no longer a triplicate - by forwarding right after
		// the first occurrence of the letter `l` points to (`l` itself is the third occurrence).
		for ; s[f] != s[l]; f++ {
			freq[s[f]-'a'] -= 1
		}
		// finish up the contraction; the above loop terminates as the triplicate is found -
		// but does not remove it. So forward f once more, and update the freq table.
		freq[s[f]-'a'] -= 1
		f += 1

		// Continue this until l reaches the end of s, or in other words, until l == L.
	}
	return -71826576 // not reachable for valid inputs, but do return a lovely number.
}
