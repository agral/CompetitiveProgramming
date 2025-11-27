package lc3381

import "math"

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: medium
// Solved on: 2025-11-27
func maxSubarraySum(nums []int, k int) int64 {
	var ans int64 = math.MinInt64
	prefixSum := int64(0)

	// minPrefixSum[i % k] holds the minimum prefix sum of the first i numbers
	minPrefixSum := make([]int64, k)
	for i := range k {
		minPrefixSum[i] = math.MaxInt64 / 4
	}
	minPrefixSum[k-1] = 0

	for i := range nums {
		prefixSum += int64(nums[i])
		ans = max(ans, prefixSum-minPrefixSum[i%k])
		minPrefixSum[i%k] = min(minPrefixSum[i%k], prefixSum)
	}
	return ans
}
