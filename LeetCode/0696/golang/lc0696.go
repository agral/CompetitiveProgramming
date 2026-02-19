package lc0696

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2026-02-19
func countBinarySubstrings(s string) int {
	L := len(s)
	currLength, lastLength := 1, 0
	currCh := s[0]
	ans := 0
	for i := 1; i < L; i++ {
		if s[i] == currCh {
			currLength += 1
		} else {
			ans += min(currLength, lastLength)
			lastLength = currLength
			currLength = 1
			currCh = s[i]
		}
	}
	ans += min(currLength, lastLength)
	return ans
}
