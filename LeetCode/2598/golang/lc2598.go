package lc2598

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: medium.
func findSmallestInteger(nums []int, value int) int {
	LEN_NUMS := len(nums)
	count := make(map[int]int)
	for _, num := range nums {
		k := num % value
		count[(k+value)%value] += 1
	}
	for i := range LEN_NUMS {
		k := i % value
		if count[k] == 0 {
			return i
		}
		count[k] -= 1
	}
	return LEN_NUMS
}
