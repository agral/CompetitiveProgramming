package lc0172

// Runtime complexity: O(log5(n)) (that's right! :-))
// Auxiliary space: O(1)
// Subjective level: medium
// Solved on: 2026-01-26
func trailingZeroes(n int) int {
	ans := 0
	for n > 0 {
		ans += n / 5
		n /= 5
	}
	return ans
}
