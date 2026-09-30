package lc1111

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: Don't bother, it's an easy question described in a badly convoluted way.
// Solved on: 2026-09-30
func maxDepthAfterSplit(seq string) []int {
	L := len(seq)
	ans := make([]int, L)

	// Represents the current nesting level.
	level := 0

	// All odd-level parens go to group "0", and all even-level parens go to group "1".
	for i, ch := range seq {
		if ch == '(' {
			level += 1
			ans[i] = (level + 1) % 2
		} else {
			ans[i] = (level + 1) % 2
			level -= 1
		}
	}
	return ans
}
