package lc3702

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2026-08-15
func longestSubsequence(nums []int) int {
	L := len(nums)
	isAllZeroes := true
	totalXor := 0
	for _, num := range nums {
		totalXor ^= num
		if num != 0 {
			isAllZeroes = false
		}
	}

	if isAllZeroes {
		// A corner case. Any subsequence (after excluding any number of zeroes)
		// still has XOR sum of 0, since it's all zeroes. Return 0.
		return 0
	}
	if totalXor == 0 {
		// then it is enough to not include any non-zero entry in the entire array -
		// and the subsequence's total xor is guaranteed to be nonzero. So return L-1.
		return L - 1
	}
	// Otherwise, the entire array forms a sequence with non-zero XOR sum.
	return L
}
