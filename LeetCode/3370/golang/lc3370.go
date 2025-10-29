package lc3370

// Runtime complexity: O(10) == O(1).
// Auxiliary space: O(10) == O(1).
// Subjective level: easy.
func smallestNumber(n int) int {
	VALID := []int{1, 3, 7, 15, 31, 63, 127, 255, 511, 1023}
	for _, num := range VALID {
		if num >= n {
			return num
		}
	}
	return -42
}
