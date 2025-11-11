package lc0474

// Runtime complexity: don't know, but lower bound is O(numZeroes * numOnes * len(strs))
// Auxiliary space: O(numZeroes * numOnes)
// Subjective level: hard.
// Solved on: 2025-11-11
func findMaxForm(strs []string, numZeroes int, numOnes int) int {
	// dp[z][o] holds the maximum size of the subset with at most `z` zeroes and `o` ones.
	dp := make([][]int, numZeroes+1)
	for z := range numZeroes + 1 {
		dp[z] = make([]int, numOnes+1)
	}

	for _, str := range strs {
		// Counts the zeroes and the ones manually. Don't want to use strings.Count just for that.
		zeroes := 0
		ones := 0
		for _, ch := range str {
			if ch == '0' {
				zeroes += 1
			} else {
				ones += 1
			}
		}

		for z := numZeroes; z >= zeroes; z-- {
			for o := numOnes; o >= ones; o-- {
				dp[z][o] = max(dp[z][o], dp[z-zeroes][o-ones]+1)
			}
		}
	}

	return dp[numZeroes][numOnes]
}
