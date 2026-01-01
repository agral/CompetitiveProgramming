package lc0066

// Runtime complexity: O(n), where n is the length of digits, `L` in code below.
// Auxiliary space: O(n) for the case of digits consisting entirely of `9`s.
// this is required by the language - have to preallocate the entire slice for
// the answer containing `1` followed by that many `0`s, cannot reuse the input slice.
// for the general case of at least one digit < 9, O(1) as the input slice is reused.
// Subjective level: easy.
// Solved on: 2026-01-01
func plusOne(digits []int) []int {
	L := len(digits)
	// 1. Check if `digits` contain only `9`s.
	isOnlyNines := true
	for i := 0; i < L && isOnlyNines; i++ {
		if digits[i] != 9 {
			isOnlyNines = false
		}
	}

	// 2a. If so, the answer is `1` followed by `L` `0`s.
	if isOnlyNines {
		ans := make([]int, L+1)
		ans[0] = 1
		return ans
	}
	// 2b. Otherwise, the answer fits in the original array.
	// Modify it by performing the standard long addition (of 1),
	// then return it.
	carry := 1
	for i := L - 1; i >= 0 && carry == 1; i-- {
		digits[i] += carry
		if digits[i] == 10 {
			digits[i] = 0
			// carry remains `1`, to be used again in the next step.
		} else {
			carry = 0
			// this completes the long addition.
		}
	}
	return digits
}
