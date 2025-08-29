package lc3021

// Runtime complexity: O(1) :-)
// Auxiliary space: O(1)
// Subjective level: easy (but the terrible description makes this problem ~medium)
func flowerGame(n int, m int) int64 {
	// x+y has to be even for Alice to win. That's it.
	// when X is even, valid Ys form the sequence 1, 3, 5, ..., m (or m-1, depending on m).
	// when X is odd, valid Ys form the sequence 2, 4, 6, ..., m-1 (or m).

	// Count of even & odd numbers present in the sequence 1, 2, ..., k, for k == either n or m,
	// already in the required int64 form:
	xEven, yEven := int64(n/2), int64(m/2)
	xOdd, yOdd := int64((n+1)/2), int64((m+1)/2)
	return xEven*yOdd + xOdd*yEven
}
