package lc2369

func validPartition(nums []int) bool {
	length := len(nums)
	// dp[i] is true iff there exists a valid partitioning for the first i numbers in nums.
	dp := make([]bool, length+1)
	dp[0] = true
	if nums[0] == nums[1] { // 2 <= len(nums) <= 10^5; so nums[0] and nums[1] are guaranteed to exist.
		dp[2] = true
	}

	for i := 3; i <= length; i++ {
		dp[i] = (dp[i-2] && (nums[i-2] == nums[i-1])) ||
			(dp[i-3] && (((nums[i-3] == nums[i-2]) && (nums[i-2] == nums[i-1])) ||
				((nums[i-3] == nums[i-2]-1) && (nums[i-2] == nums[i-1]-1))))
	}

	return dp[length]
}
