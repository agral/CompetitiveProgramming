package lc2185

import "strings"

func prefixCount(words []string, pref string) int {
	// It's basically 3042/Count Prefix and Suffix Pairs I,
	// but ignoring the suffix part.
	ans := 0
	for _, w := range words {
		if strings.HasPrefix(w, pref) {
			ans += 1
		}
	}

	return ans
}
