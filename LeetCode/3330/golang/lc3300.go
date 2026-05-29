package lc3300

// Runtime complexity: O(n)
// Auxiliary space: O(1)
func possibleStringCount(word string) int {
	ans := 1
	for i := 1; i < len(word); i++ {
		if word[i-1] == word[i] {
			ans++
		}
	}
	return ans
}
