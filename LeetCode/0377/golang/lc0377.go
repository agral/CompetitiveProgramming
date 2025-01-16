package lc0377

// Note: it is in fact a permutation sum, not a combination sum.
// (in combinations the order of selected elements does not matter,
// but in this challenge the order matters, making it a permutation sum).
// https://en.wikipedia.org/wiki/Combination
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1 // a boundary condition: there's exactly one way to make a sum of 0 - by an empty set.
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
