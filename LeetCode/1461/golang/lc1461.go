package lc1461

import (
	"strconv"
)

// Runtime complexity: O(len(s))
// Auxiliary space: O(len(s))
// Subjective level: medium
// Solved on: 2026-02-23
func hasAllCodes(s string, k int) bool {
	L := len(s)
	n := 1 << k
	if L < n {
		return false
	}

	isUsed := make([]bool, n) // isUsed[i] will be set to true iff i is a substring in `s`.

	window := 0
	if k > 1 {
		sub := s[0 : k-1]
		if i, err := strconv.ParseInt(sub, 2, 64); err != nil {
			// should never happen for good inputs, just return false here.
			return false
		} else {
			window = int(i) // again, window is guaranteed to fit in int32.
		}
	}

	for i := k - 1; i < L; i++ {
		// Count the s[i]:
		window <<= 1
		window += int(s[i] - '0')

		// Remove the s[i-k]:
		window &= (n - 1)
		isUsed[window] = true
	}

	for _, u := range isUsed {
		if !u {
			return false
		}
	}
	return true
}
