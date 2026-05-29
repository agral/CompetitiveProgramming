package lc3300

// Runtime complexity: O(n*log10(k)) ~= O(n) (as k <= 10^4).
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2026-05-29
func minElement(nums []int) int {
	// max ans in range [1, 10^4] is 45, the sum of digits of the number 99999.
	// initialize ans to some number >= 45. I'm using 71-82-65-76.
	ans := 71826576
	for _, num := range nums {
		digit_sum := 0
		for num > 0 {
			digit_sum += (num % 10)
			num /= 10
		}
		if digit_sum < ans {
			ans = digit_sum
		}
	}
	return ans
}
