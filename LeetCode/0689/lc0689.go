package lc0689

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	ans := []int{-1, -1, -1}
	sz := len(nums) - k + 1

	// sums[i] = sum(nums[i] .. nums[i+k])
	sums := make([]int, sz)

	// left[i] holds the index in [0..i) that has maximum subarray sum (sums[i])
	left := make([]int, sz)

	// right[i] holds the index in (i..sz) that has maximum subarray sum (sums[i]).
	right := make([]int, sz)

	windowSum := 0
	for i, num := range nums {
		windowSum += num
		if i >= k-1 {
			if i >= k {
				windowSum -= nums[i-k]
			}
			sums[i-k+1] = windowSum
		}
	}

	maxLeftIdx := 0
	maxRightIdx := sz - 1
	for i := 0; i < sz; i++ {
		if sums[i] > sums[maxLeftIdx] {
			maxLeftIdx = i
		}
		left[i] = maxLeftIdx
	}

	for i := sz - 1; i >= 0; i-- {
		if sums[i] >= sums[maxRightIdx] {
			maxRightIdx = i
		}
		right[i] = maxRightIdx
	}

	for i := k; i < sz-k; i++ {
		l := left[i-k]
		r := right[i+k]
		if (ans[0] == -1) ||
			(sums[ans[0]]+sums[ans[1]]+sums[ans[2]] < sums[l]+sums[i]+sums[r]) {
			ans[0] = l
			ans[1] = i
			ans[2] = r
		}
	}

	return ans
}
