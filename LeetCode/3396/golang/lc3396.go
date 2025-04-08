package lc3396

import (
	"math"
)

// Runtime complexity: O(n)
// Aux space complexity: O(n) [worst case; when all nums are different]
func minimumOperations(nums []int) int {
	// Approach: instead of deleting, start by knowing that removing *all* the elements
	// results in a distinct (empty) array. Then find the first chunk (1, 2 or 3 elements)
	// that when added results in a non-distinct array. This will be easy to check with sets.
	// Given that the last chunk may be incomplete (for arrays of size not a multiple of 3),
	// num_chunks should include only full chunks, without the last one counted separately.

	// Note: the array is guaranteed to not be empty, so optimized for that.
	// This code will panic when an empty array is provided.
	num_chunks := int(math.Ceil(float64(len(nums))/3)) - 1
	len_last_chunk := len(nums) % 3
	if len_last_chunk == 0 {
		len_last_chunk = 3
	}

	set := make(map[int]bool)
	// process the last chunk; it can have 3, 2 or 1 items.
	for n := 3 * num_chunks; n < len(nums); n++ {
		if _, isPresent := set[nums[n]]; isPresent {
			return num_chunks + 1
		}
		set[nums[n]] = true
	}
	for c := num_chunks - 1; c >= 0; c-- {
		for n := 3 * c; n < 3*c+3; n++ {
			if _, isPresent := set[nums[n]]; isPresent {
				return c + 1
			}
			set[nums[n]] = true
		}
	}
	// No duplicate found in the entire nums array. No need to delete anything.
	return 0
}
