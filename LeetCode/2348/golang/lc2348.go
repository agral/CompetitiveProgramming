package lc2348

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy
func zeroFilledSubarray(nums []int) int64 {
	// n consecutive zeroes form a total of sum(1, 2, ..., n) == n(n+1)/2 zero-filled subarrays.
	ans := int64(0)
	zeroCount := int64(0)
	for i := range nums {
		if nums[i] == 0 {
			zeroCount += 1
		} else {
			if zeroCount > 0 {
				ans += zeroCount * (zeroCount + 1) / 2
			}
			zeroCount = 0
		}
	}
	// Handle the corner case of the array ending with 0:
	if zeroCount > 0 {
		ans += zeroCount * (zeroCount + 1) / 2
	}
	return ans
}
