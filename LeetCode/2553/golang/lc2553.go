package lc2553

// Runtime complexity: O(n*log10(k)), where n is the length of nums,
// and k the average size of the length of a number in nums (in decimal digits).
// It could be argued that as 1 <= nums[i] <= 10^5, k <= 5, so that'd be O(n).

// Auxiliary space: O(|digits|) == O(n * log10(k)).
// Subjective level: easy
// Solved on: 2026-05-11
func separateDigits(nums []int) []int {
	ans := []int{}
	for _, num := range nums {
		digits := []int{}
		for num > 0 {
			digit := num % 10
			digits = append(digits, digit)
			num /= 10
		}
		// digits has the digits of num, but in low-to-high order.
		// transcribe these to ans, but in the proper high-to-low order.
		for d := len(digits) - 1; d >= 0; d-- {
			ans = append(ans, digits[d])
		}
	}
	return ans
}
