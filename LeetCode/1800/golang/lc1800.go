package lc1800

func maxAscendingSum(nums []int) int {
	ans := nums[0]
	runningSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			runningSum += nums[i]
		} else {
			runningSum = nums[i]
		}
		ans = max(ans, runningSum)
	}
	return ans
}
