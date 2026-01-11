package lc0085

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

func (s *Stack) Top() int {
	sz := len(*s)
	if sz == 0 {
		return -42 // only ever used in conjunction with IsEmpty(). Will never hit.
	}
	return (*s)[sz-1]
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Runtime complexity: O(H*W)
// Auxiliary space: O(W)
// Subjective level: hard
// Solved on: 2026-01-11
func maximalRectangle(matrix [][]byte) int {
	//H := len(matrix)
	W := len(matrix[0])
	histogram := make([]int, W)

	getLargestRectangleArea := func() int {
		area := 0
		stack := Stack{}
		for h := 0; h <= len(histogram); h++ {
			for !stack.IsEmpty() && (h == len(histogram) || histogram[stack.Top()] > histogram[h]) {
				height := histogram[stack.Top()]
				stack.Pop()
				width := 0
				if stack.IsEmpty() {
					width = h
				} else {
					width = h - stack.Top() - 1
				}
				area = max(area, height*width)
			}
			stack.Push(h)
		}
		return area
	}

	ans := 0
	for _, row := range matrix {
		for w := range W {
			if row[w] == '0' {
				histogram[w] = 0
			} else {
				histogram[w] = histogram[w] + 1
			}
		}
		ans = max(ans, getLargestRectangleArea())
	}
	return ans
}
