package lc1784

// Runtime complexity: O(n), where n is the length of the binary string.
// Auxiliary space complexity: O(1)
// Subjective level: easy
// Solved on: 2026-03-06
func checkOnesSegment(s string) bool {
	// Note: this solution does not require using strings package
	// (otherwise it's just a plain strings.Contains(s, "01") )
	// So let's roll a manual search for clusters of ones:

	// the first character of the string is guaranteed to be '1'.
	i := 0
	L := len(s)

	// 1. Skip to the first '0', or to the end of the string,
	// whichever comes first:
	for ; i < L && s[i] == '1'; i++ {
	}

	// 2. Skip to the next '1', or to the end of the string,
	// whichever comes first:
	for ; i < L && s[i] == '0'; i++ {
	}

	// 3. If i points to the end of the string, there was just one
	// segment of ones in total. Otherwise i points to the first '1'
	// starting the second cluster of ones.
	return i == L
}
