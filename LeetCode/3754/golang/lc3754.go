package lc3754

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

func (s *Stack) Len() int {
	return len(*s)
}

// Runtime complexity: O(log10(n))
// Auxiliary space: O(log10(n)) - for holding the digits on a stack.
// Subjective level: easy+
// Solved on: 2026-07-07
func sumAndMultiply(n int) int64 {
	digitSum := 0
	// these digits will come in the order from least to most significant.
	// with a stack I'm going to reverse these.
	digitStack := Stack{}

	for k := n; k > 0; k /= 10 {
		digit := k % 10
		if digit > 0 {
			digitSum += digit
			digitStack.Push(digit)
		}
	}

	newNumberWithoutZeroes := 0
	for range digitStack.Len() {
		digit, _ := digitStack.Pop()
		newNumberWithoutZeroes = 10*(newNumberWithoutZeroes) + digit
	}

	return int64(digitSum) * int64(newNumberWithoutZeroes)
}
