package lc0678

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy+
// Solved on: 2026-10-04
func checkValidString(s string) bool {
	lowLeftParenCountBoundary := 0
	highLeftParenCountBoundary := 0
	for _, ch := range s {
		switch ch {
		case '(':
			// it's definitely a '('.
			lowLeftParenCountBoundary += 1
			highLeftParenCountBoundary += 1
		case ')':
			// decrement the min count, but don't go below 0.
			// This corresponds to some stars being empty strings or an invalid string,
			// but can't tell the difference at this point yet.
			if lowLeftParenCountBoundary > 0 {
				lowLeftParenCountBoundary -= 1
			}
			// and decrement the max count. If this value becomes negative,
			// the entire string is invalid.
			highLeftParenCountBoundary -= 1
			if highLeftParenCountBoundary < 0 {
				return false
			}
		case '*':
			// decrement the lower bound, clamping to 0 for negatives.
			// (so the following treats '*' as either ')' or '':
			if lowLeftParenCountBoundary > 0 {
				lowLeftParenCountBoundary -= 1
			}

			// and increment the high bound. This handles star-to-'(' exchange.
			highLeftParenCountBoundary += 1
		}
	}
	// upon reaching the end of string: it's valid iff low boundary on left paren count equals 0.
	return lowLeftParenCountBoundary == 0
}
