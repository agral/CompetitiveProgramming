package lc1541

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy+
// Solved on: 2026-10-09
func minInsertions(s string) int {
	numRequiredRight := 0
	countInsertedLeft := 0  // holds the required and not present '(' chars, that have to be inserted manually.
	countInsertedRight := 0 // holds the required and not present ')' chars, which also did not have a corresponding '(' char. I might be overcomplicating this, a cleaner solution might exist.
	for _, ch := range s {
		if ch == '(' {
			if numRequiredRight&1 == 1 { // so if e.g. "()(*", this holds.
				countInsertedRight += 1
				numRequiredRight -= 1
			}
			numRequiredRight += 2
		} else {
			numRequiredRight -= 1     // ch is either a '(' or a ')', so at this point it's ')' guaranteed.
			if numRequiredRight < 0 { // the case of e.g. ")*", "())))))*", and so on.
				countInsertedLeft += 1
				numRequiredRight += 2
			}
		}
	}

	return countInsertedLeft + countInsertedRight + numRequiredRight

	/* WA #01
	ans := 0
	i := 0
	for i < len(s) {
		if s[i] == '(' {
			countUnmatchedLeftParens += 1
			i += 1
		} else if s[i] == ')' {
			if i+1 < L && s[i+1] == ')' {
				if countUnmatchedLeftParens > 0 {
					countUnmatchedLeftParens -= 1
				} else {
					ans += 1 // insert a '(' at the beginning of string. This balances these two "))".
				}
				i += 2 // skip the two "))" chars that we've just processed.
			} else {
				// the case of '()(*', or "*()END"; insert a ')' right after the first ')'.
				ans += 1
				countUnmatchedLeftParens -= 1
				i += 1
			}
		}
	}
	return ans + 2*numRequiredRight
	*/
}
