package lc3867

import "slices"

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Runtime complexity: ?. Surely more than O(n), but lte O(n^2).
// Auxiliary space: O(n)
// Subjective level: hard
// Solved on: 2026-09-21
func gcdSum(nums []int) int64 {
	L := len(nums)
	prefixGcd := make([]int, L)

	maxSoFar := 0
	for idx, val := range nums {
		if val > maxSoFar {
			maxSoFar = val
		}

		prefixGcd[idx] = gcd(val, maxSoFar)
	}

	slices.Sort(prefixGcd)
	ans := int64(0)
	for idx := 0; idx < L/2; idx++ { // note: this might ignore the middle element, and that's expected.
		ans += int64(gcd(prefixGcd[idx], prefixGcd[L-1-idx]))
	}

	return ans
}
