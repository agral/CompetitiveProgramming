package lc2011

// Runtime complexity: O(n) (when treating string comparison of three-letter strings as O(3) == O(1)).
// Auxiliary space: O(1).
// Subjective level: easy.
func finalValueAfterOperations(operations []string) int {
	ans := 0
	for _, op := range operations {
		if (op == "++X") || (op == "X++") {
			ans += 1
		} else {
			// op is guaranteed to be either "X--" or "--X".
			ans -= 1
		}
	}
	return ans
}
