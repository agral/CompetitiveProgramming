package lc2270

func waysToSplitArray(nums []int) int {
	sz := len(nums)
	prefixSum := make([]int, sz)
	prefixSum[0] = nums[0]
	for i := 1; i < sz; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	//total := prefixSum[sz-1]

	ans := 0
	for i := 1; i < sz; i++ {
		/*
			sum_left := prefixSum[i-1]
			sum_right := total - sum_left
			if sum_left >= sum_right {
				ans++
			}
		*/

		// optimization: don't calculate the right part's sum; it will be less than left's if
		// the left part's sum is equal to or more than half of the total.
		// In order to not divide integers, the same holds if the left sum doubled
		// is greater than or equal to the total sum.

		if 2*prefixSum[i-1] >= prefixSum[sz-1] {
			ans++
		}
	}
	return ans
}
