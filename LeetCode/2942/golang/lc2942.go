package lc2942

import "strings"

// Runtime complexity: O(n * |word[i]|); generally O(n^2).
// Auxiliary space: O(1)
func findWordsContaining(words []string, x byte) []int {
	ans := []int{}
	xstr := string(x)
	for i := range words {
		if strings.Contains(words[i], xstr) {
			ans = append(ans, i)
		}
	}
	return ans
}
