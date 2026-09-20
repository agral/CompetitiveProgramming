package lc3498

// Runtime complexity: O(n)
// Auxiliary space: O(1)
// Subjective level: easy
// Solved on: 2026-09-20
func reverseDegree(s string) int {
	ans := 0
	for i, ch := range s {
		reversedOneIndexed := int('z'-ch) + 1 // so 'z' - 'a' + 1 is 26, 'z' - 'b' + 1 is 25, 'z' - 'z' + 1 is 1.
		ans += (i + 1) * reversedOneIndexed
	}
	return ans
}
