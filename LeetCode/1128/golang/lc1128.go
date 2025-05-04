package lc1128

// Runtime complexity: O(n)
// Auxiliary space: O(100) == O(1).
func numEquivDominoPairs(dominoes [][]int) int {
	hashes := make(map[int]int)
	for _, domino := range dominoes {
		var h int
		if domino[0] < domino[1] {
			h = 10*domino[0] + domino[1]
		} else {
			h = 10*domino[1] + domino[0]
		}
		hashes[h] += 1
	}
	ans := 0
	for _, h := range hashes {
		if h > 0 {
			ans += h * (h - 1) / 2
		}
	}
	return ans
}
