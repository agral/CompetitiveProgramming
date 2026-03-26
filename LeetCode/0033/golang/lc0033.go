package lc0033

// Runtime complexity: O(log(L))
// Auxiliary space: O(1)
// Subjective level: medium-
// Solved on: 2026-03-26
func search(nums []int, target int) int {
	// 1. Start the search with full range in consideration:
	L := len(nums)
	left := 0
	right := L - 1

	// 2. Divide the range in half. Determine whether each partition
	// contains a pivot (it does when subrange_left > subrange_right).
	for left < right {
		mid := left + (right-left)/2
		if nums[left] <= nums[mid] {
			// handle the case of left partition not containing a pivot point.
			// does it contain the target?
			if nums[left] <= target && target <= nums[mid] {
				// the target is contained within the left partition exclusively;
				// continue searching there.
				right = mid
			} else {
				// the target is contained within the right partition exclusively
				// (the one with a pivot point). Continue searching there.
				left = mid + 1
			}
		} else {
			// handle the case of left partition containing a pivot point;
			// the code is similar to the one handling pivot point in the right partition.
			// Does the pivot-free right partition contain the target?
			if nums[mid+1] <= target && target <= nums[right] {
				// the target is contained within the right partition exclusively;
				// continue searching there.
				left = mid + 1
			} else {
				// the target is contained within the left partition (the one with
				// a pivot point). Continue searching there.
				right = mid
			}
		}
	}

	// left now either points to the searched for target;
	// or target is entirely not present in nums.
	if nums[left] == target {
		return left
	}
	return -1
}
