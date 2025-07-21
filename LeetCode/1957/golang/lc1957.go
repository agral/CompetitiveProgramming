package lc1957

// Runtime complexity: O(n)
// Auxiliary space: O(n) (for storing the modified "fancy" string)
func makeFancyString(s string) string {
	len_s := len(s)
	if len_s < 3 {
		return s
	}
	ans := make([]byte, 0, len_s)
	// Copy the first two chars:
	ans = append(ans, s[0])
	ans = append(ans, s[1])
	var prev byte = s[0]
	var last byte = s[1]

	for i := 2; i < len_s; i++ {
		if s[i] != prev || s[i] != last {
			ans = append(ans, s[i])
		}
		prev = last
		last = s[i]
	}
	return string(ans)
}
