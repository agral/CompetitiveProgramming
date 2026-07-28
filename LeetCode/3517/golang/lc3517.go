package lc3517

import "slices"

// Runtime complexity: O(sort)
// Auxiliary space: O(n)
// Subjective level: medium-, or even an easy++.
// Solved on: 2026-07-28
func smallestPalindrome(s string) string {
	L := len(s)
	half := []byte(s[0 : L/2]) // half is a slice of bytes, not a string. So it can be sorted.
	mid := ""                  // also take care of the odd-one-out middle char in inputs of odd length.
	if L%2 == 1 {
		mid = string(s[L/2])
	}
	slices.Sort(half)
	revHalf := make([]byte, L/2)
	copy(revHalf, half)
	slices.Reverse(revHalf)

	return string(half) + string(mid) + string(revHalf)
}
