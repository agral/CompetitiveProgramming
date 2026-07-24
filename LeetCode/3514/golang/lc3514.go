package lc3514

import "slices"

// Runtime complexity: O(M^2), where M is the maximum number in nums
// Auxiliary space: O(M)
// Subjective level: medium+.
// Solved on: 2026-07-24
func uniqueXorTriplets(nums []int) int {
	// the absolute total possible maximum number resulting from A xor B xor C,
	// where A, B, C are in nums - it's twice the maximum of nums.
	max := 2 * slices.Max(nums)

	aXorB := make([]bool, max)
	for _, a := range nums {
		for _, b := range nums {
			aXorB[a^b] = true
		}
	}

	aXorBXorC := make([]bool, max)
	// x := a xor b
	for x := range max {
		if aXorB[x] {
			for _, c := range nums {
				aXorBXorC[x^c] = true
			}
		}
	}

	ans := 0
	for _, xorValue := range aXorBXorC {
		if xorValue {
			ans += 1
		}
	}
	return ans
}
