package lc3364

import "math"

// Runtime complexity: O(n^2)
// Auxiliary space: O(1)
// Subjective level: medium-
// Solved on: 2026-11-14
func minimumSumSubarray(nums []int, l int, r int) int {
	L := len(nums)
	ans := math.MaxInt32
	for windowLength := l; windowLength <= r; windowLength++ {
		// Calculate the initial sum (of the leftmost window of length windowLength):
		sum := 0
		for k := 0; k < windowLength; k++ {
			sum += nums[k]
		}
		// Check if this initial sum is greater than zero; if so, it might be an answer:
		if sum > 0 {
			ans = min(ans, sum)
		}

		// Now iterate over all the remaining windows of length windowLength.
		// To get the next window's sum, remove the value from the left,
		// and add the value from the right.
		for w := windowLength; w < L; w++ {
			sum -= nums[w-windowLength]
			sum += nums[w]
			if sum > 0 {
				ans = min(ans, sum)
			}
		}
	}
	// At this point all the possible windows have been checked.
	// If ans is still a MaxInt32, there were no positive sum windows, so return -1.
	if ans == math.MaxInt32 {
		return -1
	}
	// Otherwise return the answer, which is the minimum sum of all the windows.
	return ans
}
