package lc2302

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// both are optimal for this problem.
func countSubarrays(nums []int, k int64) int64 {
	ans := int64(0)
	sum := int64(0)
	for l, r := 0, 0; r < len(nums); r++ {
		sum += int64(nums[r])
		for sum*int64(r-l+1) >= k {
			sum -= int64(nums[l])
			l += 1
		}
		ans += int64(r - l + 1)
	}
	return ans
}
