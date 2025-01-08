package lc3042

import "strings"

func countPrefixSuffixPairs(words []string) int {
	ans := 0
	for p := 0; p < len(words)-1; p++ {
		for w := p + 1; w < len(words); w++ {
			if strings.HasPrefix(words[w], words[p]) && strings.HasSuffix(words[w], words[p]) {
				ans += 1
			}
		}
	}
	return ans
}
