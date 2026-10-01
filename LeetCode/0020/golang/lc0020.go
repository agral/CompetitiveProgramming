package lc0020

type CharStack []byte

func (s *CharStack) Push(val byte) {
	*s = append(*s, val)
}

func (s *CharStack) Pop() byte {
	sz := len(*s)
	if sz == 0 {
		// should throw an error, but I'm only interested in matching the parens.
		// So just return a predefined value that is not a left paren.
		// 'E', short for "Error, stack empty" works well.
		return 'E'
	}
	val := (*s)[sz-1]
	*s = (*s)[:sz-1]
	return val
}

func (s *CharStack) Len() int {
	return len(*s)
}

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: easy
// Solved on: 2026-10-01
func isValid(s string) bool {
	stack := CharStack{}
	for _, ch := range s {
		bch := byte(ch) // golang has its bad parts, runes vs bytes is one such bad part.
		switch bch {
		case '(', '{', '[':
			stack.Push(bch)
		case ')', '}', ']':
			corresponding := stack.Pop()
			if (bch == ')' && corresponding != '(') ||
				(bch == '}' && corresponding != '{') ||
				(bch == ']' && corresponding != '[') {
				return false
			}
		}
	}
	// reached the end of string without nonmatching parentheses - means the string is valid.
	// return true

	// Nope, WA #01 - you have to have an empty stack as well.
	// Otherwise, we're approving invalid strings such as "{", "({[" and so on.
	return stack.Len() == 0
}
