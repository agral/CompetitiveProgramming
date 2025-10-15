package lc3350

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy
func maxIncreasingSubarrays(nums []int) int {
	lengthIncreasing := 1
	lengthPreviousIncreasing := 0
	ans := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			lengthIncreasing += 1
		} else {
			lengthPreviousIncreasing = lengthIncreasing
			lengthIncreasing = 1
		}
		ans = max(ans, lengthIncreasing/2, min(lengthIncreasing, lengthPreviousIncreasing))
	}
	return ans
}
