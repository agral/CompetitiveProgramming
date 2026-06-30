package lc1358

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: medium. Has several options for optimizing it. Liked it.
// Solved on: 2026-06-30
func numberOfSubstrings(s string) int {
	lastIndexOfA := -1
	lastIndexOfB := -1
	lastIndexOfC := -1
	ans := 0
	for i, ch := range s {
		switch ch {
		case 'a':
			lastIndexOfA = i
		case 'b':
			lastIndexOfB = i
		case 'c':
			lastIndexOfC = i
		}
		ans += min(lastIndexOfA, lastIndexOfB, lastIndexOfC) + 1
	}
	return ans
}
