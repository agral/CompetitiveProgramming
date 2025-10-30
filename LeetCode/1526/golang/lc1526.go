package lc1526

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy
func minNumberOperations(target []int) int {
	ans := target[0]
	for i := 0; i < len(target)-1; i++ {
		if target[i] < target[i+1] {
			ans += target[i+1] - target[i]
		}
	}
	return ans
}
