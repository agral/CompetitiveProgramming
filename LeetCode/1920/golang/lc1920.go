package lc1920

// Runtime complexity: O(n)
// Auxiliary space: O(1)
func buildArray(nums []int) []int {
	length := len(nums)
	ans := make([]int, length)
	for i := range length {
		ans[i] = nums[nums[i]]
	}
	return ans
}
