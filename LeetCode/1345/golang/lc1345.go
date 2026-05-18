package lc1345

import "fmt"

type Queue []int

func (q *Queue) Enqueue(val int) {
	*q = append(*q, val)
}

func (q *Queue) Dequeue() (int, error) {
	if len(*q) == 0 {
		return 0, fmt.Errorf("Dequeue() from an empty Queue.")
	}
	val := (*q)[0]
	*q = (*q)[1:]
	return val, nil
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}

func (q *Queue) Len() int {
	return len(*q)
}

// Runtime complexity: O(n)
// Auxiliary space: O(n)
// Subjective level: medium+
// Solved on: 2026-05-18
func minJumps(arr []int) int {
	L := len(arr)
	ans := L - 1 // in worst case, there's no shortcuts and one has to
	// perform L-1 jumps to get from index 0 to index L-1.

	graph := make(map[int][]int)
	seen := make(map[int]bool, L)
	q := Queue{}
	q.Enqueue(0)
	seen[0] = true

	for i := range L {
		graph[arr[i]] = append(graph[arr[i]], i)
	}

	for stepNum := 0; !q.Empty(); stepNum++ {
		for range q.Len() {
			cellIdx, _ := q.Dequeue()
			if cellIdx == L-1 {
				return stepNum
			}
			seen[cellIdx] = true
			cellValue := arr[cellIdx]

			// one can jump to the right from every cell except the last one (so, when cellIdx < L-1).
			if cellIdx < L-1 {
				graph[cellValue] = append(graph[cellValue], cellIdx+1)
			}
			// one can jump to the left from every cell except the first one (so, when cellIdx > 0).
			if cellIdx > 0 {
				graph[cellValue] = append(graph[cellValue], cellIdx-1)
			}

			for _, idx := range graph[cellValue] {
				if !seen[idx] { // process each cell only once.
					q.Enqueue(idx)
				}
			}

			graph[cellValue] = []int{}
		}
	}

	return ans
}
