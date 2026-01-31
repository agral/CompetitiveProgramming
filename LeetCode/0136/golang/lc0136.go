package lc0136

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2026-01-31
func singleNumber(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}
