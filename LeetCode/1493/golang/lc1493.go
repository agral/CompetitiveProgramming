package lc1493

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: Medium
func longestSubarray(nums []int) int {
	l := 0
	numZeroes := 0
	for _, num := range nums {
		if num == 0 {
			numZeroes += 1
		}
		if numZeroes > 1 {
			if nums[l] == 0 {
				numZeroes -= 1
			}
			l += 1
		}
	}
	return len(nums) - l - 1
}
