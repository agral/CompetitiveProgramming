package lc2411

// Runtime complexity: O(30n) == O(n)
// Auxiliary space: O(n)
func smallestSubarrays(nums []int) []int {
	const MAX_SET_BIT int = 30
	aux := make([]int, MAX_SET_BIT)
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		for j := range MAX_SET_BIT {
			if (nums[i]>>j)&1 == 1 {
				aux[j] = i
			}
			ans[i] = max(ans[i], aux[j]-i+1)
		}
	}
	return ans
}
