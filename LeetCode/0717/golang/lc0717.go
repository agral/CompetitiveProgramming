package lc0717

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2025-11-18
func isOneBitCharacter(bits []int) bool {
	pos := 0
	L := len(bits)
	for pos < L-1 {
		if bits[pos] == 1 {
			pos += 2
		} else {
			pos += 1
		}
		if pos == L-1 {
			return true
		}
	}
	return pos == L-1
}
