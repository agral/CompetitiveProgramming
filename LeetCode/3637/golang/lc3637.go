package lc3637

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2026-02-03
func isTrionic(nums []int) bool {
	// first pair must be strictly increasing.
	if nums[1] <= nums[0] {
		return false
	}
	phase := 1 // phases: up-down-up, encoded as 1-2-3. We're in the first "up" phase now.
	for i := 2; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return false // repeated numbers mean broken "strictly decreasing/increasing" condition.
		}
		if phase == 1 && nums[i] < nums[i-1] {
			// found first strictly decreasing sequence; switch to phase 2.
			phase = 2
		} else if phase == 2 && nums[i] > nums[i-1] {
			// found the second strictly increasing sequence; switch to phase 3.
			phase = 3
		} else if phase == 3 && nums[i] < nums[i-1] {
			// another strictly decreasing sequence means that this slice is not trionic.
			return false
		}
	}
	return phase == 3
}
