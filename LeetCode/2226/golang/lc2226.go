package lc2226

// Runtime complexity: O(n * log(sumCandies)), where n is the length of `candies` slice.
// (this n stems from the requirement of recalculating the number of satisfied children for
// every new midpoint. Each such recalculation is done in O(n)).
// Aux space complexity: O(1).
func maximumCandies(candies []int, k int64) int {
	getNumSatisfiedChildren := func(midpoint int) int64 {
		var ans int64 = 0
		for _, candy := range candies {
			ans += int64(candy / midpoint)
		}
		return ans
	}

	left := 1
	var sumCandies int64 = 0
	for _, candy := range candies {
		sumCandies += int64(candy)
	}
	right := int(sumCandies / k)

	for left < right {
		var mid int = (left + right) / 2
		if getNumSatisfiedChildren(mid) < k {
			right = mid
		} else {
			left = mid + 1
		}
	}

	if getNumSatisfiedChildren(left) < k {
		left -= 1
	}
	return left
}
