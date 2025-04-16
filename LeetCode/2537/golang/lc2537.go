package lc2537

// Runtime complexity: O(n).
// Aux space complexity: O(n).
func countGood(nums []int, k int) int64 {
	var ans int64 = 0
	pairs := 0
	count := make(map[int]int)
	l := 0

	for r := range nums {
		pairs += count[nums[r]]
		count[nums[r]] += 1
		for pairs >= k {
			count[nums[l]] -= 1
			pairs -= count[nums[l]]
			l += 1
		}
		ans += int64(l)
	}

	return ans
}
