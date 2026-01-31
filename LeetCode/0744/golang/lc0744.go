package lc0744

// Runtime complexity: O(log(n))
// Auxiliary space: O(1)
// Subjective level: medium
// Solved on: 2026-01-31
func nextGreatestLetter(letters []byte, target byte) byte {
	// perform a binary search for the next letter:
	l := 0
	// h := len(letters) - 1 note: no, have to overshoot with h=len(letters),
	// to handle the case of the answer not being present in the array!
	// then in the end: l=h=len(letters), which is 0 modulo len(letters).
	// otherwise would have to check again when this algo terminates whether
	// the answer is correct.
	h := len(letters)
	for l < h {
		mid := (l + h) / 2
		if letters[mid] <= target {
			l = mid + 1
		} else {
			h = mid
		}
	}
	return letters[l%len(letters)]
}
