package lc2913

// Runtime complexity: O(n^2)
// Auxiliary space: O(n)
// Subjective level: easy+
// Solved on: 2026-10-15
func sumCounts(nums []int) int {
	const MAX_NUM = 100
	L := len(nums)
	ans := 0
	for b := 0; b < L; b++ { // b: holds the beginning index of the currently analyzed subarray
		set := make([]int, MAX_NUM+1) // zero-initializes all the entries.
		num_uniques := 0
		for e := b; e < L; e++ { // e: holds the ending (last) index of the currently analyzed subarray.
			set[nums[e]] += 1
			if set[nums[e]] == 1 { // just discovered a new number in this subarray
				num_uniques += 1
			}
			ans += num_uniques * num_uniques
		}
	}
	return ans
}
