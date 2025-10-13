package lc2273

import (
	"slices"
	"sort"
	"strings"
)

// Runtime complexity: O(n^2)
// Auxiliary space: O(n)
// Subjective level: medium
func removeAnagrams(words []string) []string {
	len_words := len(words)
	sorted := make([]string, len_words)
	for i, word := range words {
		runes := strings.Split(word, "")
		sort.Strings(runes)
		s := strings.Join(runes, "")
		sorted[i] = s
	}

	for len_words > 1 {
		isChanged := false
		for i := 1; i < len_words; i++ {
			if sorted[i-1] == sorted[i] {
				words = slices.Delete(words, i, i+1)
				sorted = slices.Delete(sorted, i, i+1)
				len_words -= 1
				isChanged = true
				i -= 1
			}
		}
		if !isChanged {
			return words
		}
	}
	return words
}
