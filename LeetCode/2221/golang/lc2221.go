package lc2221

// Runtime complexity: O(n^2)
// Auxiliary space: O(1) (reusing the nums array)
// Subjective level: easy
func triangularSum(nums []int) int {
	LEN_NUMS := len(nums)
	for i := LEN_NUMS - 1; i > 0; i-- {
		for j := 1; j <= i; j++ {
			nums[j-1] = (nums[j-1] + nums[j]) % 10
		}
	}
	return nums[0]
}
