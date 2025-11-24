package lc1018

// Runtime complexity: O(n)
// Auxiliary space: O(n) (required for the answer array; otherwise O(1))
// Subjective level: easy
// Solved on: 2025-11-24
func prefixesDivBy5(nums []int) []bool {
	N := len(nums)
	ans := make([]bool, N)
	num := 0
	for i, bitValue := range nums {
		num = (2*num + bitValue) % 5
		ans[i] = (num == 0)
	}
	return ans
}
