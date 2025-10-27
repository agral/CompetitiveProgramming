package lc2125

import "strings"

// Runtime complexity: O(m*n); where m and n are the dimensions of `bank` 2D matrix.
// Auxiliary space: O(1).
// Subjective level: easy.
func numberOfBeams(bank []string) int {
	prev := 0
	ans := 0
	for _, row := range bank {
		curr := strings.Count(row, "1")
		if curr > 0 {
			ans += (curr * prev)
			prev = curr
		}
	}
	return ans
}
