package lc0394

import "strings"

type StringInt struct {
	s string
	i int
}

type Stack []StringInt

func (s Stack) Push(val StringInt) Stack {
	return append(s, val)
}
func (s Stack) Pop() (Stack, StringInt) {
	sz := len(s)
	if sz == 0 {
		panic("Just tried to Pop() from an empty Stack.")
	}
	return s[:sz-1], s[sz-1]
}

func decodeString(s string) string {
	var stack Stack
	ans := ""
	number := 0

	for _, c := range s {
		if c >= '0' && c <= '9' {
			number *= 10
			number += int(c - '0')
		} else if c == '[' {
			stack = stack.Push(StringInt{ans, number})
			number = 0
			ans = ""
		} else if c == ']' {
			var si StringInt
			stack, si = stack.Pop()
			ans = si.s + strings.Repeat(ans, si.i)
		} else {
			ans += string(c)
		}
	}

	return ans
}
