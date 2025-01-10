package lc0916

func wordSubsets(words1 []string, words2 []string) []string {
	const NUM_LETTERS = 'z' - 'a' + 1

	getStringFrequency := func(input string) []int {
		freq := make([]int, NUM_LETTERS)
		for _, ch := range input {
			freq[ch-'a'] += 1
		}
		return freq
	}

	isSufficient := func(freq []int, countAll2 []int) bool {
		for i := 0; i < NUM_LETTERS; i++ {
			if countAll2[i] > freq[i] {
				return false
			}
		}
		return true
	}

	// Compute the union of all words2's frequencies:
	countAll2 := make([]int, NUM_LETTERS)
	for _, word := range words2 {
		freq := getStringFrequency(word)
		for i := 0; i < NUM_LETTERS; i++ {
			countAll2[i] = max(countAll2[i], freq[i])
		}
	}

	ans := []string{}
	for _, word := range words1 {
		if isSufficient(getStringFrequency(word), countAll2) {
			ans = append(ans, word)
		}
	}

	return ans
}
