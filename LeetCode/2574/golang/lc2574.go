package lc2574

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Runtime complexity: O(n)
// Auxiliary space: O(n) (for the answer, excluding that it's O(1)).
// Subjective level: easy+.
// Solved on: 2026-06-06
func leftRightDifference(nums []int) []int {
	L := len(nums)
	// there is no slices.Sum()...
	sum := 0
	for _, val := range nums {
		sum += val
	}
	leftSum := 0
	rightSum := sum - nums[0]
	ans := make([]int, L)
	ans[0] = rightSum // as leftSum is 0, and rightSum is (sum-nums[0]).
	for i := 1; i < L-1; i++ {
		leftSum += nums[i-1]
		rightSum -= nums[i]
		ans[i] = abs(leftSum - rightSum)
	}
	ans[L-1] = sum - nums[L-1] // as left is sum - nums[L-1], and right is 0.
	return ans
}
