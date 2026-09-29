package lc1717

import "slices"

type Stack []byte

func (s *Stack) Push(val byte) {
	*s = append(*s, val)
}

func (s *Stack) Pop() byte { // just don't call it on an empty stack.
	sz := len(*s) // every call preceded by an IsEmpty() invocation.
	val := (*s)[sz-1]
	*s = (*s)[:sz-1]
	return val
}

func (s *Stack) Peek() byte { // just don't call it on an empty stack.
	return (*s)[len(*s)-1] // return the last element without popping it.
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Reverse() {
	slices.Reverse(*s)
}

// Runtime complexity: O(len(s)) == O(n)
// Auxiliary space: O(len(s)) == O(n)
// Subjective level: medium+ / hard-. Pretty convoluted, requires using two stacks,
// one of which needs to be treated as a FIFO queue at some point.
// Maybe this could have been done more cleverly... But for me, a harder medium, or an easier hard.
// Solved on: 2026-09-29
func maximumGain(s string, x int, y int) int {
	//fmt.Println("---")
	getGain := func(strExpensive string, valExpensive int, strCheap string, valCheap int) int {
		ans := 0
		stackExpensive := Stack{}

		// first try removing `strExpensive` for `valExpensive` score gain, where possible:
		for _, ch := range s {
			bch := byte(ch) // they're runes when iterating over a string with the handy for.
			if !stackExpensive.IsEmpty() && stackExpensive.Peek() == strExpensive[0] &&
				bch == strExpensive[1] {
				ans += valExpensive
				stackExpensive.Pop()
				//fmt.Println("Expensive!")
			} else {
				stackExpensive.Push(bch)
			}
		}

		// now, try removing `strCheap` from whatever is left after previous removal
		// (the contents of `stackExpensive`):
		stackCheap := Stack{}
		// Note: stackExpensive needs to be reversed, so that the contents will be served
		// as if it was a queue (i.e. a FIFO manner, not a LIFO as with a stack).
		stackExpensive.Reverse()
		for !stackExpensive.IsEmpty() {
			bch := stackExpensive.Pop()
			if !stackCheap.IsEmpty() && stackCheap.Peek() == strCheap[0] &&
				bch == strCheap[1] {
				ans += valCheap
				stackCheap.Pop()
				//fmt.Println("Cheap!")
			} else {
				stackCheap.Push(bch)
			}
		}
		return ans
	}

	// Prefer removing either "ab" or "ba", whichever results in higher score gain:
	if x >= y {
		return getGain("ab", x, "ba", y)
	} else {
		return getGain("ba", y, "ab", x)
	}
}
