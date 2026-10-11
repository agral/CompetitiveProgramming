package lc2778

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2026-10-11
func sumOfSquares(nums []int) int {
	n := len(nums)
	ans := 0
	for i, num := range nums {
		if n%(i+1) == 0 {
			ans += num * num
		}
	}
	return ans
}
