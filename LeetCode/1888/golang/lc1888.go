package lc1888

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: medium
// Solved on: 2026-03-07
func minFlips(s string) int {
	// Note: repeated trick not needed; worked out a solution that
	// doesn't depend on `Type 1` alterations at all.
	// repeated := s + s
	L := len(s)

	// counts[0] := count of even indexed zeroes
	// counts[1] := count of odd indexed zeroes
	// counts[2] := count of even indexed ones
	// counts[3] := count of odd indexed ones
	counts := []int{0, 0, 0, 0}
	for i, ch := range s {
		num := int(ch - '0') // will be either 0 or 1, depending on the input character
		counts[2*num+(i%2)] += 1
	}

	// At every step the answer is the minimum of:
	// a) make all the zeroes in even indices -> ones, and all the ones in odd indices -> zeroes
	// b) make all the ones in even indices -> zeroes, and all the zeroes in odd indices -> ones
	ans := min(counts[0]+counts[3], counts[2]+counts[1])

	// simulate moving the 0, 0..1, 0..2, ..., 0..nth characters to the end of the string:
	for i, ch := range s { // could have used s[:L-1]; but that might be actually slower.
		num := int(ch - '0')
		counts[2*num+(i%2)] -= 1
		counts[2*num+((i+L)%2)] += 1
		ans = min(ans, counts[0]+counts[3], counts[2]+counts[1])
	}

	return ans
}
