package lc2433

// Runtime complexity: O(n)
// Auxiliary space: O(1) (O(n) if you have to count the answer array)
// Subjective level: easy
// Solved on: 2026-01-31
func findArray(pref []int) []int {
	L := len(pref)
	ans := make([]int, L)
	cumulativeXor := 0
	for i := range pref {
		ans[i] = cumulativeXor ^ pref[i]
		cumulativeXor ^= ans[i]
	}
	return ans
}
