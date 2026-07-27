package lc0628

import "slices"

// Runtime complexity: O(sort) == O(nlogn)
// Auxiliary space: O(sort) == O(n)
// Subjective level: easy
// Solved on: 2026-07-26
func maximumProduct(nums []int) int {
	//bestPos1, bestPos2, bestPos3 := 0, 0, 0 // will hold the three biggest positive integers in nums
	//bestNeg1, bestNeg2 := 0, 0              // will hold the two biggest (absolute value) negative integers in nums.
	// Or zeroes - if there are not enough positive/negative ints.
	// no, don't. Too much effort. Just sort, pick three largest items:
	// a) either the three largest nums ([L-1], [L-2], [L-3]), or
	// b) the two lowest nums (i.e. two max absolute value negative numbers), and one biggest num.

	// So, step one, sort:
	slices.Sort(nums)

	// Step two, answer. It's guaranteed that the length of nums is at least three:
	L := len(nums)
	return max(nums[L-1]*nums[L-2]*nums[L-3], nums[0]*nums[1]*nums[L-1])
}
