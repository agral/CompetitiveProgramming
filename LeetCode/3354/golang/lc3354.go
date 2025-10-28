package lc3354

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: easy
func countValidSelections(nums []int) int {
	n := len(nums)
	prefixSum := make([]int, n) // prefixSum[i] = sum(nums[0] to nums[i-1]).
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}
	suffixSum := make([]int, n) // suffixSum[i] = sum(nums[i+1] to nums[n-1]).
	for i := n - 2; i >= 0; i-- {
		suffixSum[i] = suffixSum[i+1] + nums[i+1]
	}
	ans := 0

	for i := range n {
		if nums[i] == 0 {
			absDiff := abs(prefixSum[i] - suffixSum[i])
			switch absDiff {
			case 1:
				ans += 1 // can go from here to the directian with a larger sum; will succeed.
			case 0:
				// in other words, if prefixSum[i] == suffixSum[i].
				// this means that one can go from here in any direction (left/right);
				// will succeed in both cases.
				ans += 2
			}
		}
	}

	return ans
}
