package lc3174

func clearDigits(s string) string {
	stack := []rune{} // a nice poor man's LIFO stack implementation on a slice of runes.
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			length := len(stack)     // performs a "pop" of a previous character.
			stack = stack[:length-1] // note: only non-digits are stored on the stack.
		} else {
			stack = append(stack, ch) // "push"-es a non-digit rune to the LIFO stack.
		}
	}
	return string(stack)
}
