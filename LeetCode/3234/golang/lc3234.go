package lc3234

// Runtime complexity: O(n * sqrt(n))
// Auxiliary space: O(1)
// Subjective level: hard
// Solved on: 2025-11-15
func numberOfSubstrings(s string) int {
	N := len(s)
	ans := 0

	for zero := 0; zero*zero+zero <= N; zero++ {
		lastBadPos := -1
		numZeroes := 0
		numOnes := 0
		l := 0
		for r := 0; r < N; r++ {
			// Update the count with current right-side char:
			if s[r] == '0' {
				numZeroes++
			} else {
				numOnes++
			}

			// Shrink current window if possible:
			for l < r {
				if s[l] == '0' && numZeroes > zero {
					// remove an extraneous zero:
					lastBadPos = l
					numZeroes -= 1
					l++
				} else if s[l] == '1' && numOnes-1 >= zero*zero {
					// remove an extraneous one:
					numOnes -= 1
					l++
				} else {
					// did not remove neither '0' nor '1' this loop;
					// so no more characters can be removed with this `r`.
					break
				}
			}

			if numZeroes == zero && numOnes >= zero*zero {
				// Current window holds at least one dominant ones sequence.
				ans += (l - lastBadPos)
			}
		}
	}
	return ans
}
