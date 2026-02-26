package lc1404

import (
	"strconv"
)

// Runtime complexity: O(len(s))
// Auxiliary space: O(1)
// Subjective level: medium-
// Solved on: 2026-02-26
func numSteps(s string) int {
	ans := 0
	bytes := []byte(s)

	// pop the "0" characters at the end:
	for bytes[len(bytes)-1] == '0' {
		ans += 1
		bytes = bytes[:len(bytes)-1]
	}

	// If the original string consisted of '1' followed by '0's, the number becomes odd only at the end
	// when n == 1. This is the only case when the number remains even for the entire time.
	if len(bytes) == 1 {
		return ans
	}

	// the number is now odd, so:
	// - adding 1 to bytes costs exactly 1 step
	// - all the '1' characters become '0' in 1 step each;
	// - all the '0' characters become '1' in one extra step (adding "1");
	//   then become '0' again in 1 step each, so in summary:
	// - so in summary: it costs 2 to remove each '0', 1 to remove each '1', with one extra '1' cost.
	for _, ch := range bytes {
		if ch == '0' {
			ans += 2
		} else {
			ans += 1
		}
	}
	return 1 + ans
}

// It could have been so easy
func oldNumSteps(s string) int {
	n, _ := strconv.ParseInt(s, 2, 64)
	ans := 0
	for n > 1 {
		ans += 1
		if n&1 == 0 {
			n >>= 1
		} else {
			n += 1
		}
	}
	return ans
}
