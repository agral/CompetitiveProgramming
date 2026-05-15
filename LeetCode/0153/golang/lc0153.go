package lc0153

// Runtime complexity: O(log(n))
// Auxiliary space: O(1)
// Subjective level: medium+
// Solved on: 2026-05-15
func findMin(nums []int) int {
	L := len(nums)
	// the minimum value is somewhere between left and right.
	left := 0
	right := L - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}
