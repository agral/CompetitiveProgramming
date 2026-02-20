package lc0761

import (
	"sort"
	"strings"
)

// Runtime complexity: O(n^2)
// Auxiliary space: O(n)
// Subjective level: hard
// Solved on: 2026-02-20
func makeLargestSpecial(s string) string {
	L := len(s)
	specials := []string{}
	balance := 0

	i := 0 // holds the beginning of the last known special string
	for l := 0; l < L; l++ {
		if s[l] == '1' {
			balance += 1
		} else {
			balance -= 1
		}

		if balance == 0 {
			innerStr := s[i+1 : l]
			specials = append(specials, "1"+makeLargestSpecial(innerStr)+"0")
			i = l + 1
		}
	}

	sort.Sort(sort.Reverse(sort.StringSlice(specials)))

	return strings.Join(specials, "")
}
