package lc1464

// Runtime complexity: O(N)
// Auxiliary space: O(1)
// Subjective level: trivially easy
// Solved on: 2026-07-27
func maxProduct(nums []int) int {
	// Choose the indices of two biggest numbers. That's it.
	// This strategy wouldn't work for negative numbers, when at least two
	// such numbers are present (and then would only require a second check),
	// but `nums` will only contain entries in range [1, 1000].

	biggest, semiBiggest := 0, 1
	if nums[0] < nums[1] {
		biggest, semiBiggest = 1, 0
	}
	for i := 2; i < len(nums); i++ {
		if nums[i] > nums[biggest] {
			semiBiggest = biggest
			biggest = i
		} else if nums[i] > nums[semiBiggest] {
			semiBiggest = i
		}
	}
	return (nums[biggest] - 1) * (nums[semiBiggest] - 1)
}
