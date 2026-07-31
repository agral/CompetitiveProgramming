package lc3016

import "slices"

// Runtime complexity: O(sort) == O(nlogn)
// Auxiliary space: O(sort)
// Subjective level: easy+
// Solved on: 2026-07-30
// Note: The solution is same as 3014.
func minimumPushes(word string) int {
	// 1. Count the number of times each individual letter appears in word:
	NUM_LETTERS := 'z' - 'a' + 1
	count := make([]int, NUM_LETTERS)
	for i := 0; i < len(word); i++ {
		count[word[i]-'a'] += 1
	}

	// 2. Sort these in descending order:
	slices.SortFunc(count, func(lhs, rhs int) int {
		return rhs - lhs
	})

	// 3. Calculate the cost: there are 8 "programmable" keys.
	// The first 8 most expensive letters get to be programmed first as keys 2-9.
	// These will require 1 click to type.
	// The next 8 most expensive letters get to be programmed second as keys 2-9.
	// Each of these will require 2 clicks each to type. And so on.
	// So, the total amount of typing the entire `word` is calculated as:
	ans := 0
	for i := range count {
		ans += count[i] * (1 + (i / 8))
	}
	return ans
}
