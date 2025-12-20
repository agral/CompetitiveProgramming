package lc0944

// Runtime complexity: O(n * wordlen)
// Auxiliary space: O(1)
// Subjective level: easy, but weird; this tests... nothing? Unnecessarily convoluted?
// Solved on: 2025-12-20
func minDeletionSize(strs []string) int {
	n := len(strs)
	wordlen := len(strs[0])
	ans := 0

	for x := range wordlen {
		for y := range n - 1 {
			if strs[y][x] > strs[y+1][x] {
				ans += 1
				break
			}
		}
	}
	return ans
}
