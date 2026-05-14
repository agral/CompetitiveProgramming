package lc2784

// Runtime complexity: O(n)
// Auxiliary space: O(n) (since n <= 100, it could be argued that it's O(1) for both runtime and space).
// Subjective level: easy.
// Solved on: 2026-05-14
func isGood(nums []int) bool {
	// Note: constraints say that:
	// * 1 <= nums.length <= 100
	// * 1 <= num[i] <= 200
	// but that means that if any number n > 99 is seen, the array is necessarily invalid!
	// (as base[99] == []int{1, 2, 3, ..., 98, 99, 99} is the largest array still in the constraints;
	// with length of exactly 100.

	// Approach: prepare a slice of 100 ints (99 will be used, from 1 to 99, 0 indexed).
	// Count the occurences of all the numbers in nums.
	// Verify that there's exactly 1 of each number <= n, and exactly 2 of n.
	count := [100]int{}
	n := -1
	for _, num := range nums {
		if num >= 100 {
			return false
		}
		count[num] += 1
		n = max(n, num)
	}
	if count[n] != 2 {
		return false
	}
	for i := 1; i < n; i++ {
		if count[i] != 1 {
			return false
		}
	}
	return true
}
