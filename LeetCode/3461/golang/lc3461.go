package lc3461

// Runtime complexity: O(n^2)
// Auxiliary space: O(n)
// Subjective level: easy
func hasSameDigits(s string) bool {
	LEN := len(s)
	digits := make([]int, LEN)
	for i, ch := range s {
		digits[i] = int(ch - '0')
	}

	for o := LEN - 1; o >= 2; o-- {
		for i := 0; i < o; i++ {
			digits[i] = (digits[i] + digits[i+1]) % 10
		}
	}

	return digits[0] == digits[1]
}
