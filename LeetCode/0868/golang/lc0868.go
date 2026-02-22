package lc0868

// Runtime complexity: O(log2(n))
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2026-02-22
func binaryGap(n int) int {
	const SET_BIT_NOT_SEEN_YET int = -1
	ans := 0
	prev := -1
	for n > 0 {
		if n&1 == 1 {
			if prev == SET_BIT_NOT_SEEN_YET {
				prev = 0
			} else {
				ans = max(ans, prev)
				prev = 0
			}
		}
		if prev != SET_BIT_NOT_SEEN_YET {
			prev += 1
		}
		n >>= 1
	}
	return ans
}
