package lc2762

func continuousSubarrays(nums []int) int64 {
	ans := int64(1) // that will be nums[0]. nums is guaranteed to have at least one element.
	left := nums[0] - 2
	right := nums[0] + 2
	left_idx := 0

	for right_idx := 1; right_idx < len(nums); right_idx++ {
		if left <= nums[right_idx] && nums[right_idx] <= right {
			left = max(left, nums[right_idx]-2)
			right = min(right, nums[right_idx]+2)
		} else {
			// nums[right_idx] does not fit in currently considered sliding window.
			// Construct the new window with only right_idx:
			left = nums[right_idx] - 2
			right = nums[right_idx] + 2
			left_idx = right_idx
			for nums[right_idx]-2 <= nums[left_idx] && nums[left_idx] <= nums[right_idx]+2 { // for is Go's while
				left = max(left, nums[left_idx]-2)
				right = min(right, nums[left_idx]+2)
				left_idx -= 1
			}
			left_idx += 1
		}
		ans += int64(right_idx - left_idx + 1)
	}
	return ans
}
