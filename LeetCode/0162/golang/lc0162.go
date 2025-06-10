package lc0162

// Runtime complexity: O(log(SZ))
// Auxiliary space: O(1)
func findPeakElement(nums []int) int {
	SZ := len(nums)

	// the peak is somewhere between low and high.
	low := 0
	high := SZ - 1
	for low < high {
		mid := (low + high) / 2
		if nums[mid] > nums[mid+1] {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}
