package lc1190

import "fmt"

type Stack []int

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}

func (s *Stack) Pop() (int, error) {
	sz := len(*s)
	if sz == 0 {
		return 0, fmt.Errorf("Pop() from an empty Stack.")
	}
	val := (*s)[sz-1]
	*s = (*s)[:sz-1]
	return val, nil
}

// Runtime complexity: O(N^2)
// Auxiliary space: O(N)
// Subjective level: medium
// Solved on: 2026-09-28
// By: agral
func reverseParentheses(s string) string {
	//func reverseParenthesesBruteforce(s string) string {
	ans := []rune{}
	stack := Stack{}

	// Reverses the indicated part of the rune slice, in-place.
	reverse := func(b, e int) {
		for ; b < e; b, e = b+1, e-1 {
			ans[b], ans[e] = ans[e], ans[b]
		}
	}

	for _, ch := range s {
		if ch == '(' {
			stack.Push(len(ans))
		} else if ch == ')' {
			offset, _ := stack.Pop()
			reverse(offset, len(ans)-1)
		} else { // a normal char:
			ans = append(ans, ch)
		}
	}

	return string(ans)
}

// This can be done in O(N) -- I've already solved that in C++ in O(N).
// But it's kind of tedious to write in Golang. Maybe I'll return to this one
// in the future. For now, happy with an O(N^2) bruteforce soln.
