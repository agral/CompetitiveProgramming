package lc2434

import "fmt"

// Runtime complexity: O(n)
// Auxiliary space: O(26*n) == O(n)
// Subjective level: medium. Also annoying: had to roll out my own char stack data structure.
func robotWithString(s string) string {
	const NUM_CHARS = 'z' - 'a' + 1
	count := make([]int, NUM_CHARS)
	for _, ch := range s {
		count[ch-'a'] += 1
	}
	getLowestAvailableChar := func() byte {
		for i := range NUM_CHARS {
			if count[i] > 0 {
				return byte('a' + i)
			}
		}
		return byte('a') // should never be reached, but I can't be bothered to add proper error handling.
	}

	ans := []byte{}

	stack := ByteStack{}
	for _, ch := range s {
		stack.Push(byte(ch))
		count[ch-'a'] -= 1
		lowestAvailChar := getLowestAvailableChar()
		for !stack.Empty() && stack.Top() <= lowestAvailChar {
			char, _ := stack.Pop()
			ans = append(ans, char)
		}
	}

	for !stack.Empty() {
		char, _ := stack.Pop()
		ans = append(ans, char)
	}

	return string(ans)
}

type ByteStack []byte

func (s *ByteStack) Empty() bool {
	return len(*s) == 0
}

func (s *ByteStack) Push(val byte) {
	*s = append(*s, val)
}

func (s *ByteStack) Pop() (byte, error) {
	sz := len(*s)
	if sz == 0 {
		return 0, fmt.Errorf("Pop() from an empty ByteStack.")
	}
	val := (*s)[sz-1]
	*s = (*s)[:sz-1]
	return val, nil
}

func (s *ByteStack) Top() byte {
	sz := len(*s)
	if sz == 0 {
		panic("Top() called on an empty ByteStack")
	}
	return (*s)[sz-1]
}
