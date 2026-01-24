package lc0009

// Runtime complexity: O(log10(x))
// Auxiliary space: O(1)
// Subjective level: medium.
// Note: easy to do it with to-string-and-back conversion, but it's specifically not allowed here.
// Solved on: 2026-01-24
func isPalindrome(x int) bool {
	if x < 0 {
		// any number with "-" in the beginning can not be a palindrome.
		return false
	}

	reversed := int64(0)
	x64 := int64(x)

	for x64 > 0 {
		reversed = 10*reversed + x64%10
		x64 /= 10
	}
	return reversed == int64(x)
}
