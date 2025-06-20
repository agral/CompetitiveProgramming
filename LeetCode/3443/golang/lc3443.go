package lc3443

import "strings"

// Runtime complexity: O(4n) == O(n)
// Auxiliary space: O(n)
func maxDistance(s string, k int) int {
	mutate := func(direction string) int {
		ans := 0
		pos := 0
		opposite := 0
		for _, ch := range s {
			if strings.ContainsRune(direction, ch) {
				pos++
			} else {
				pos--
				opposite++
			}
			ans = max(ans, pos+2*min(k, opposite))
		}
		return ans
	}

	return max(mutate("NE"), mutate("NW"), mutate("SW"), mutate("SE"))
}
