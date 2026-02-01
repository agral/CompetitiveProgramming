package lc3010

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2026-02-01
func minimumCost(nums []int) int {
	// Have to divide the slice into three slices so that the first elements' sum
	// is minimal. The first element has to be nums[0] no matter what.
	// So the only thing left to do is to find two smallest elements in the rest
	// of the slice. It is guaranteed that the input slice `nums` contains
	// at least three elements, and no more than fifty.
	// Could be solved by sorting nums[1:] and taking first two values,
	// but that's O(sort) == O(nlogn); instead rolling out a simple custom search
	// with O(n) complexity:
	second, third := nums[1], nums[2]
	// keep the invariant second >= third:
	if second < third {
		second, third = third, second
	}

	for i := 3; i < len(nums); i++ {
		if nums[i] < second {
			if nums[i] < third {
				//  ith number is lower than both second and third.
				// Store hitherto `third` value as `second`, and store ith number as `third`.
				second = third
				third = nums[i]
			} else {
				// ith number is lower than `second`, but not lower than `third`.
				// keep third unchanged, store nums[i] as `second`.
				second = nums[i]
			}
		}
	}
	return nums[0] + second + third
}
