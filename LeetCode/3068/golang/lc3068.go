package lc3068

import "math"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Runtime complexity: O(n)
// Auxiliary space: O(1)
func maximumValueSum(nums []int, k int, edges [][]int) int64 {
	// Note: this solution purposefully ignores the `edges` parameter.
	optimalSum := int64(0)
	numChangedItems := 0
	minChangedDiff := math.MaxInt32

	for _, num := range nums {
		numXorK := num ^ k
		optimalSum += int64(max(num, numXorK))
		if num^k > num {
			numChangedItems++
		}
		changeDiff := abs(num - numXorK)
		minChangedDiff = min(minChangedDiff, changeDiff)
	}

	if numChangedItems%2 == 0 {
		return optimalSum
	}
	return optimalSum - int64(minChangedDiff)
}
