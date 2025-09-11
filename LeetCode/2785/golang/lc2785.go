package lc2785

import "slices"

// Runtime complexity: O(sort) == O(n*log(n))
// Auxiliary space: O(n)
// Subjective level: medium
func sortVowels(s string) string {
	isVowel := func(r rune) bool {
		switch r {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			return true
		}
		return false
	}
	vowels := []rune{}
	for _, letter := range s {
		if isVowel(letter) {
			vowels = append(vowels, letter)
		}
	}

	slices.Sort(vowels)

	ans := make([]rune, len(s))
	currVowelIdx := 0
	for i, letter := range s {
		if isVowel(letter) {
			ans[i] = vowels[currVowelIdx]
			currVowelIdx += 1
		} else {
			ans[i] = rune(s[i])
		}
	}
	return string(ans)
}
