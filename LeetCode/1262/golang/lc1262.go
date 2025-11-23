package lc1262

import "slices"

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: medium
// Solved on: 2025-11-23
func maxSumDivThree(nums []int) int {
	sums := make([]int, 3) // sums[i] == max sum seen so far, such that this sum (mod 3) == i
	for _, num := range nums {
		cloneSums := slices.Clone(sums)
		for _, sum := range cloneSums {
			sums[(sum+num)%3] = max(sums[(sum+num)%3], sum+num)
		}
	}
	return sums[0]
}
