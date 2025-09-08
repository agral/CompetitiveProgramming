package lc1317

// works correctly only for positive integers;
// as only positive integers will be used in the challenge.
func hasZeroDigitInDecimalRepresentation(n int) bool {
	for n > 0 {
		digit := n % 10
		if digit == 0 {
			return true
		}
		n /= 10
	}
	return false
}

// Runtime complexity: O(n * log10(n));
// - hasZeroDigitInDecimalRepresentation iterates over the number of digits, which is ~log10(n)
// - and the entire algo runs over up to (n/2) numbers in range [1, n/2].
// Auxiliary space: O(1)
// Subjective level: medium
func getNoZeroIntegers(n int) []int {
	for a := 1; a <= n/2; a++ {
		b := n - a
		if !(hasZeroDigitInDecimalRepresentation(a) || hasZeroDigitInDecimalRepresentation(b)) {
			return []int{a, b}
		}
	}
	// guaranteed unreachable; don't want to panic (runtime overhead), but have to return something.
	//panic("Could not solve.")
	return []int{-1, -1}
}
