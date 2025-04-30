package lc1295

// Runtime complexity: O(1)
// Auxiliary space: O(1)
func findNumbers(nums []int) int {
	ans := 0
	for _, num := range nums {
		// Each num is guaranteed to be in [1, 10^5] range.
		// So this does not even require logarithmic calculations; a simple conditional is sufficient.
		if (num >= 10 && num <= 99) || (num >= 1000 && num <= 9999) || num == 100_000 {
			ans++
		}
	}
	return ans
}
