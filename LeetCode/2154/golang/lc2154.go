package lc2154

// Runtime complexity: O(n)
// Auxiliary space: O(n) for the map
// Subjective level: easy
// Solved on: 2025-11-19
func findFinalValue(nums []int, original int) int {
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		m[num] = true
	}

	for m[original] {
		original *= 2
	}
	return original
}
