package lc3289

// Runtime complexity: O(n)
// Auxiliary space: O(100) == O(1)
// Subjective level: easy
func getSneakyNumbers(nums []int) []int {
	ans := make([]int, 0, 2)
	seen := make([]int, 100)
	for _, num := range nums {
		seen[num] += 1
		if seen[num] == 2 {
			ans = append(ans, num)
		}
	}
	return ans
}
