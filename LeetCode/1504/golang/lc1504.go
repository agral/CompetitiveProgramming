package lc1504

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
		fmt.Println("Called Top() from an empty Stack.")
		return -1
	}
	val := (*s)[sz-1]
	return val
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Runtime complexity: O(H*W)
// Auxiliary space: O(W)
// Subjective level: medium+.
func numSubmat(mat [][]int) int {
	H := len(mat)
	W := len(mat[0])
	ans := 0
	accumulator := make([]int, W)
	for h := range H {
		for w := range W {
			if mat[h][w] == 0 {
				accumulator[w] = 0
			} else {
				accumulator[w] += 1
			}
		}
		/* ans +=  count */
		var stack Stack
		submatrices := make([]int, W)
		for w := range W {
			for !stack.IsEmpty() && accumulator[stack.Top()] >= accumulator[w] {
				stack.Pop()
			}
			if !stack.IsEmpty() {
				prevIdx := stack.Top()
				submatrices[w] = submatrices[prevIdx] + (w-prevIdx)*accumulator[w]
			} else {
				submatrices[w] = (w + 1) * accumulator[w]
			}
			stack.Push(w)
		}

		delta := 0
		for _, val := range submatrices {
			delta += val
		}
		ans += delta
	}
	return ans
}
