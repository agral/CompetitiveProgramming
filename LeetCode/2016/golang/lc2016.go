package lc2016

// Runtime complexity: O(n)
// Auxiliary space: O(1)
func maximumDifference(nums []int) int {
	minSoFar := nums[0]
	ans := 0
	for i := 1; i < len(nums); i++ {
		minSoFar = min(minSoFar, nums[i])
		ans = max(ans, nums[i]-minSoFar)
	}
	if ans == 0 {
		return -1
	}
	return ans
}
