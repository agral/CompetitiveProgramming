package lc0007

import "math"

// Runtime complexity: O(log10(x))
// Auxiliary space: O(1)
// Subjective level: medium
// Solved on: 2026-01-25
func reverse(x int) int {
	x64 := int64(x)
	ans64 := int64(0)
	for x64 != 0 {
		ans64 *= 10
		ans64 += x64 % 10
		x64 /= 10
	}
	if ans64 < math.MinInt32 || ans64 > math.MaxInt32 {
		return 0
	}
	return int(ans64)
}
