package lc3392

// Runtime complexity: O(n)
// Auxiliary space: O(1)
func countSubarrays(nums []int) int {
	ans := 0
	for i := range len(nums) - 2 {
		if nums[i+1] == 2*(nums[i]+nums[i+2]) {
			ans += 1
		}
	}
	return ans
}
