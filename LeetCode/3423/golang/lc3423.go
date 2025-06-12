package lc3423

// Runtime complexity: O(n)
// Auxiliary space: O(1)
func maxAdjacentDistance(nums []int) int {
	n := len(nums)
	ans := abs(nums[0] - nums[n-1])
	for k := 0; k < n-1; k++ {
		ans = max(ans, abs(nums[k]-nums[k+1]))
	}
	return ans
}

// Returns an absolute value of an integer.
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
