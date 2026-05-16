package lc0154

// Runtime complexity:
// - typical (average) case: O(log(n))
// - worst case (all duplicates): O(n).
// probably can't be done better, but maybe it can? Don't think so.
// Auxiliary space: O(1)
// Subjective level: medium+/hard
// Solved on: 2026-05-16
func findMin(nums []int) int {
	L := len(nums)
	// the minimum value is somewhere between left and right.
	left := 0
	right := L - 1
	for left < right {
		mid := (left + right) / 2
		// the rest is similar to 0153, but first handle the case of numbers being equal:
		if nums[mid] == nums[right] {
			right -= 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}
