package lc2243

import (
	"strconv"
	"strings"
)

// Runtime complexity: O(n)
// Auxiliary space: O(n) (for the string builder)
// Subjective level: medium. It could have been easy, but mixing strings and ints make it a medium imo.
// Solved on: 2026-11-16
func digitSum(s string, k int) string {
	for k < len(s) {
		newNumbers := strings.Builder{}
		for idxK := 0; idxK < len(s); idxK += k {
			sum := 0
			for idx := idxK; idx < min(len(s), idxK+k); idx++ {
				sum += int(s[idx] - '0')
			}
			newNumbers.WriteString(strconv.Itoa(sum))

		}
		s = newNumbers.String()
	}

	return s
}
