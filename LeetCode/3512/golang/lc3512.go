package lc3512

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2025-11-29
func minOperations(nums []int, k int) int {
	// calculate the total sum of all the elements of the nums slice:
	sum := 0
	for _, num := range nums {
		sum += num
	}

	ans := sum % k
	return ans
}
