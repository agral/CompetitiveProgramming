package lc0921

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy+
// Solved on: 2026-10-06
func minAddToMakeValid(s string) int {
	// Count left and right parentheses; return the absolute difference.
	// (this is easy since we're allowed to insert _any_ parenthesis
	// at _any_ position in the string in one move).
	countLeftParens := 0
	countRightParens := 0
	numHotfixes := 0
	for i := range s {
		if s[i] == '(' {
			countLeftParens += 1
		} else {
			countRightParens += 1
		}
		// but fix the closing-parens-before-opening-parens immediately:
		if countRightParens > countLeftParens {
			countRightParens -= 1
			numHotfixes += 1
		}
	}
	return abs(countLeftParens-countRightParens) + numHotfixes
}
