package lc1689

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy / waste of time?
// Solved on: 2026-03-01
func minPartitions(n string) int {
	ans := 0
	for _, ch := range n {
		ans = max(ans, int(ch-'0'))
	}
	return ans
}
