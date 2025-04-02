package lc2873

func maximumTripletValue(nums []int) int64 {
	ans := int64(0)
	maxRunningNum := 0  // holds max(nums[i])
	maxRunningDiff := 0 // holds max(nums[i]-nums[j]) for i, j in len(nums) and i < j
	for _, num := range nums {
		ans = max(ans, int64(maxRunningDiff)*int64(num))
		maxRunningDiff = max(maxRunningDiff, maxRunningNum-num)
		maxRunningNum = max(maxRunningNum, num)
	}
	return ans
}
