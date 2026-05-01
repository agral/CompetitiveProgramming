package lc0396

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: medium+
// Solved on: 2026-05-01
func maxRotateFunction(nums []int) int {
	// 1. Calculate the sum of the `nums` array:
	total := 0
	for _, val := range nums {
		total += val
	}

	// 2. Calculate F(0) (that is, no rotation of `nums`):
	f := 0
	for i, val := range nums {
		f += i * val
	}

	ans := f

	// 3. Check all the possible rotations; pick the one that yields highest score.
	L := len(nums)
	for k := L - 1; k > 0; k-- {
		f += total - L*nums[k]
		ans = max(ans, f)
	}
	return ans
}
