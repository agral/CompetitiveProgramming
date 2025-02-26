package lc1749

// Runtime complexity: O(n)
// Auxiliary space complexity: O(1)
func maxAbsoluteSum(nums []int) int {
	runningSum := 0
	maxPrefixSum, minPrefixSum := 0, 0
	for _, num := range nums {
		runningSum += num
		maxPrefixSum = max(maxPrefixSum, runningSum)
		minPrefixSum = min(minPrefixSum, runningSum)
	}
	return maxPrefixSum - minPrefixSum
}
