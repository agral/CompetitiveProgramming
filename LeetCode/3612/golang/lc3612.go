package lc3612

// Runtime complexity: O(n^2)
//   - as both '#' and '%' operations are themselves O(n), and there can be at most n of these.
//   - where `n` is the length of the input string `s`.
//
// Auxiliary space: O(2^n) (for holding the duplicated stack contents).
// Subjective level: medium.
// Solved on: 2026-06-16
func processStr(s string) string {
	stack := []byte{}
	for i := range s { // can't go `for _, ch := range s`, as then ch is treated as rune.
		ch := s[i] // This roundabout approach works fine, though - ch is a byte.
		switch ch {
		case '*':
			L := len(stack)
			if L > 0 {
				stack = stack[:L-1] // remove the last character from the output stack.
			}

		case '#':
			// duplicate the current stack contents, stack = stack+stack.
			// won't work: stack = append(stack[:0:0], stack...)
			// so let's manually duplicate a slice of bytes.
			L := len(stack)
			duplicated := make([]byte, 2*L)
			for i := range L {
				duplicated[i] = stack[i]
			}
			for i := range L {
				duplicated[L+i] = stack[i]
			}
			stack = duplicated

		case '%':
			// reverse the current stack contents:
			L := len(stack)
			for l := range L / 2 {
				stack[l], stack[L-1-l] = stack[L-1-l], stack[l]
			}

		default:
			// just add the character to the output stack.
			stack = append(stack, ch)
		}
	}

	return string(stack)
}
