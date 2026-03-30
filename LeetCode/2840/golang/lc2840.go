package lc2840

// Runtime complexity: O(n) == O(len(s1) == len(s2))
// Auxiliary space: O(26 * 2) == O(1)
// Subjective level: medium
// Solved on: 2026-03-30
func checkStrings(s1 string, s2 string) bool {
	NUM_LETTERS := 'z' - 'a' + 1 // 26 in total.
	counts := make([][]int, NUM_LETTERS)
	for n := range NUM_LETTERS {
		// the requirement of a difference of j-i being even means that
		// there are only two "states" to keep track of, (j-i) % 2.
		// the difference is either even or odd; no other option is possible.
		counts[n] = make([]int, 2)
	}

	n := len(s1) // == len(s2) guaranteed.
	for i := range n {
		idx1 := s1[i] - 'a' // letters of s1 and s2 are guaranteed to be lowercase
		idx2 := s2[i] - 'a' // english letters; so they are easily converted to slice indices.
		counts[idx1][i%2] += 1
		counts[idx2][i%2] -= 1
	}

	// now look for "unbalanced" letters; if any is found, the strings are not equal.
	for l := range NUM_LETTERS {
		if counts[l][0] > 0 || counts[l][1] > 0 {
			return false
		}
	}
	return true
}
