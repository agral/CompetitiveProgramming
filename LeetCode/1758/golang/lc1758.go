package lc1758

// Runtime complexity: O(len(s)) == O(n)
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2026-03-05
func minOperations(s string) int {
	numFlipsZero := 0 // numFlipsToMakeStringAlternatingStartingWithZero is a bit too Java-ey.
	numFlipsOne := 0
	for i, ch := range s {
		// for a string starting with zero, ith bit should be:
		// - set (1), if i is odd (1, 3, 5, ...);
		// - unset (0), if i is even (0, 2, 4, ...);
		// and for a string starting with one, ith bit should be:
		// - set for even-numbered bits (0, 2, 4, ..); unset for odd-numbered bits.
		if i&1 == 0 { // so if the bit is even-numbered (#0, #2, #4, ...):
			if ch == '1' {
				numFlipsZero += 1
			} else { // implicitly, ch == '0':
				numFlipsOne += 1
			}
		} else { // otherwise, for bits that are odd-numbered (#1, #3, #5, ...):
			if ch == '0' {
				numFlipsZero += 1
			} else {
				numFlipsOne += 1
			}
		}
	}
	return min(numFlipsZero, numFlipsOne)
}
