package lc0961

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: easy
// Solved on: 2026-01-02
func repeatedNTimes(nums []int) int {
	seen := map[int]bool{}
	for _, num := range nums {
		if seen[num] {
			return num
		}
		seen[num] = true
	}

	// won't ever happen on valid inputs.
	return -42
}
