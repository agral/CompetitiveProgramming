package lc3741

import "math"

// Note: This is the verbatim copy of my Golang's answer to LC3740.
// Runtime complexity: O(len(nums)). Can be argued that it's O(100) == O(1).
// Auxiliary space: O(2*len(nums)).
// Subjective level: easy+
// Solved on: 2026-04-10
func minimumDistance(nums []int) int {
	const MAX = 100
	const NO_GOOD_TUPLES = math.MaxInt32
	ans := NO_GOOD_TUPLES
	// 1. For each number, store the two indices this number was last seen at.
	// Initially, this mapping will be empty for each number:
	prevIndices := map[int][]int{}
	for i := range MAX + 1 {
		prevIndices[i] = []int{}
	}
	// 2. Process each number in nums.
	// I) If it appears for the 1st or the 2nd time, just store its index in prevIndices.
	// II) when it appears for the 3rd or further time, update the good tuple score
	//    and shift the prevIndices so that the new index is kept.
	for i, n := range nums {
		if len(prevIndices[n]) < 2 { // variant I
			prevIndices[n] = append(prevIndices[n], i)
		} else { // variant II
			// the tuple score (a, b, c) is abs(b-a) + abs(c-b) + abs(a-c).
			// Since we keep a < b < c by iterating forward, abs can be dropped:
			// score == b-a + c-b + c-a (as c > a, the last result is negative,
			// and abs makes it positive). This reduces to:
			// score == 2 * (c-a); meaning that the middle term b is unused.
			tupleScore := 2 * (i - prevIndices[n][0])
			ans = min(ans, tupleScore)
			prevIndices[n][0] = prevIndices[n][1]
			prevIndices[n][1] = i

			// also please note that with this approach it is not necessary to check *all* the
			// tuples (of which there is choose-3-n) == (n!) / (n-3)!(3!),
			// only the closest ones, exactly (n-2).
		}
	}
	if ans == NO_GOOD_TUPLES {
		return -1
	}
	return ans
}
