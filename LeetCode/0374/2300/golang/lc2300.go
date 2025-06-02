package lc2300

import "slices"

// Runtime complexity: O(sort) + O(n * log(m))
// Auxiliary space: O(n), required for storing the answer.
func successfulPairs(spells []int, potions []int, success int64) []int {
	n := len(spells)
	m := len(potions)
	slices.Sort(potions)
	ans := make([]int, n)
	for i, spell := range spells {
		lower_bound := 0
		upper_bound := m
		for lower_bound < upper_bound {
			mid := (lower_bound + upper_bound) / 2
			power := int64(spell) * int64(potions[mid])
			if power < success {
				lower_bound = mid + 1
			} else {
				upper_bound = mid
			}
		}
		// The number of potions that work are lower-bound and all the potions to the right.
		// The total count of this set is m - lower_bound:
		ans[i] = m - lower_bound
	}
	return ans
}
