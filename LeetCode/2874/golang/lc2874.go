package lc2874

// Note: this is the *same* question as lc2873.
// The description, the examples and their answers are all the same, *verbatim*.
// I was searching for a difference for some good minutes. I feel cheated now :-)
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
