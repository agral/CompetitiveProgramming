package lc3838

// Runtime complexity: O(n)
// Auxiliary space: O(len(words)) -> O(n)
// Subjective level: medium-
// Solved on: 2026-06-13
func mapWordWeights(words []string, weights []int) string {
	const NUM_LETTERS int = 'z' - 'a' + 1
	W := len(words)
	ans := make([]byte, W)
	for w := range W {
		// the weight of a word is a sum of the weights of its characters.
		// the weight of each character is to be read from the `weights` map.
		weight := 0
		for _, ch := range words[w] {
			weight += weights[ch-'a']
		}
		weight %= NUM_LETTERS

		// finally, map the resulting weight to a reverse alphabetical order;
		// where 0 -> 'z', 1 -> 'y', ..., 24 -> 'b', 25 -> 'a'.
		offset := NUM_LETTERS - weight - 1
		// -1 above is for I want 'a' to be at offset 0, not 1.
		ans[w] = byte('a' + offset)
	}
	return string(ans)
}
