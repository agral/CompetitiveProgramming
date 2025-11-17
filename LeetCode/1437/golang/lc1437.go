package lc1437

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2025-11-17
func kLengthApart(nums []int, k int) bool {
	lastOne := -(k + 1)
	for idx, val := range nums {
		if val == 1 {
			if idx-lastOne <= k {
				return false
			}
			lastOne = idx
		}
	}
	return true
}
