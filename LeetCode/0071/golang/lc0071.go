package lc0071

import (
	"strings"
)

// A stack of string entries
type Stack []string

func (s *Stack) Push(val string) {
	*s = append(*s, val)
}

func (s *Stack) Pop() string {
	sz := len(*s)
	if sz == 0 {
		return ""
	}
	val := (*s)[sz-1]
	*s = (*s)[:sz-1]
	return val
}

// Remove the last element, if possible. Don't bother returning popped values. Faster than Pop().
func (s *Stack) Del() {
	sz := len(*s)
	if sz > 0 {
		*s = (*s)[:sz-1]
	}
}

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: medium
// Solved on: 2026-01-24
func simplifyPath(path string) string {
	components := strings.Split(path, "/")
	stack := Stack{}
	for _, component := range components {
		switch component {
		case "..":
			stack.Del()
		case ".", "":
			continue
			// do nothing, this is the case of repeated "/" sign.
		default:
			stack.Push(component)
		}
	}

	// Handle the case of only "/", ".." & "." present in the absolute path
	// (these leave the stack empty):
	if len(stack) == 0 {
		return "/"
	}

	var sb strings.Builder
	for _, entry := range stack {
		sb.WriteString("/")
		sb.WriteString(entry)
	}
	ans := sb.String()
	return ans
}
