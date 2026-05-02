package lc0788

// Runtime complexity: O(log(i))
// Auxiliary space: O(1)
func isGood(i int) bool {
	// good numbers consist of at least one of [2, 5, 6, 9], any [0, 1, 8] and no [3, 4, 7] digits.
	isSymmetricDigitPresent := false
	for ; i > 0; i /= 10 {
		digit := i % 10
		switch digit {
		case 3, 4, 7:
			return false
		case 2, 5, 6, 9:
			isSymmetricDigitPresent = true
		} // else, ignore the [0 || 1 || 8] case. It changes nothing.
	}
	return isSymmetricDigitPresent
}

// Runtime complexity: O(n * log(n))
// Auxiliary space: O(1)
// Subjective level: medium
// Solved on: 2026-05-02
func rotatedDigits(n int) int {
	ans := 0
	for i := 1; i <= n; i++ {
		if isGood(i) {
			ans += 1
		}
	}
	return ans
}
