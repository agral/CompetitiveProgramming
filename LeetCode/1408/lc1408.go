package lc1408

import (
	"strings"
)

func stringMatching(words []string) []string {
	seen := map[string]bool{}
	for _, needle := range words {
		for _, haystack := range words {
			if len(needle) < len(haystack) {
				if strings.Contains(haystack, needle) {
					seen[needle] = true
					continue
				}
			}
		}
	}
	// Translate my poor man's set into the requested slice of strings:
	ans := []string{}
	for k := range seen {
		ans = append(ans, k)
	}
	return ans
}
