package lc0233

// Runtime complexity: O(log10(n))
// Auxiliary space: O(1)
// Subjective level: hard
// Solved on: 2026-01-31
func countDigitOne(n int) int {
	ans := 0

	// good luck getting that one right in an interview.
	for pow10 := 1; pow10 <= n; pow10 *= 10 {
		divisor := 10 * pow10
		quotient := n / divisor
		remainder := n % divisor
		if quotient > 0 {
			ans += pow10 * quotient
		}
		if remainder >= pow10 {
			ans += min(pow10, remainder+1-pow10)
		}
	}
	return ans
}
