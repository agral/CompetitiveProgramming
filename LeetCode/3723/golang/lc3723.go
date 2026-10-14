package lc3723

import "strings"

// Runtime complexity: O(log10(num)); generally depends on the number of decimal digits in num.
// Auxiliary space: O(log10(num))
// Subjective level: medium
// Solved on: 2026-10-14
func maxSumOfSquares(num int, sum int) string {
	if sum > 9*num {
		// then the number can't be represented using decimal digits. Return an empty string.
		return ""
	}
	numNines := sum / 9
	remainder := sum % 9
	ans := ""
	// first use as many digits 9 as possible to form the good integer.
	if numNines > 0 {
		ans += strings.Repeat("9", numNines)
	}

	// then, if `remainder` != 0, append a `remainder` digit at the end of string.
	if remainder > 0 {
		// remainder is anything between {1, 2, ..., 8}. This converts it to a string:
		ans += string(byte(remainder) + '0')
	}

	// finally, pad with zeroes to get the required digits' count. This does not change the score.
	if len(ans) < num {
		ans += strings.Repeat("0", num-len(ans))
	}

	return ans
}
