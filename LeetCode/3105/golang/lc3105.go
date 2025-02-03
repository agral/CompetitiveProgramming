package lc3105

func longestMonotonicSubarray(nums []int) int {
	currIncreasing := 1
	currDecreasing := 1
	ans := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			currIncreasing++
			currDecreasing = 1
		} else if nums[i] < nums[i-1] {
			currIncreasing = 1
			currDecreasing++
		} else {
			currIncreasing = 1
			currDecreasing = 1
		}
		ans = max(ans, currIncreasing, currDecreasing)
	}

	return ans
}
