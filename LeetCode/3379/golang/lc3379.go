package lc3379

// Runtime complexity: O(n)
// Auxiliary space: O(n) (for the answer array. Except that, constant space is used.)
// Subjective level: easy
// Solved on: 2026-02-05
func constructTransformedArray(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range nums {
		sourceIndex := (i + nums[i] + (1000 * n)) % n
		ans[i] = nums[sourceIndex]
	}
	return ans
}
