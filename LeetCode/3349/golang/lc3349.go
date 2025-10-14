package lc3349

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy
func hasIncreasingSubarrays(nums []int, k int) bool {
	lengthIncreasing := 1
	lengthPreviousIncreasing := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			lengthIncreasing += 1
		} else {
			lengthPreviousIncreasing = lengthIncreasing
			lengthIncreasing = 1
		}
		if lengthIncreasing >= 2*k || (lengthIncreasing >= k && lengthPreviousIncreasing >= k) {
			return true
		}
	}
	return false
}
